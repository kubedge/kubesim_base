#!/usr/bin/env bash
#
# bin/release.sh — cut a coordinated multi-module release of kubesim_base.
#
# kubesim_base publishes several Go modules from ONE repo (root + config/ connected/
# grpc/go/ health/ arpscan/). Go tags a sub-module as "<dir>/vX.Y.Z" and the root as
# "vX.Y.Z". Some sub-modules import others (today: grpc/go -> config), which creates a
# chicken/egg: grpc/go cannot require config@vNEW until config@vNEW is tagged AND
# pushed, and its go.sum can't be computed until the tag is fetchable.
#
# This script therefore releases in DEPENDENCY ORDER. A module is tagged only after
# every in-repo module it depends on has been: tagged, pushed, bumped in this module's
# go.mod (go get <dep>@vNEW), tidied, and the bump committed to main. Leaf modules
# carry no bump commit and are tagged at the current HEAD; dependent modules get a
# "release: bump ..." commit and are tagged at that commit.
#
# It is deliberately generic (discovers modules + edges from go.mod), so the same
# script serves the sibling sim repos, which share this multi-module shape.
#
# Usage:
#   bin/release.sh v0.1.25             # cut the release (pushes commits + tags)
#   bin/release.sh v0.1.25 --dry-run   # print the exact plan, touch nothing
#
# Preconditions: run on an up-to-date, clean `main`; `gh`/push access to origin.
#
set -euo pipefail

MODPATH_ROOT="github.com/kubedge/kubesim_base"

# ---- args ----------------------------------------------------------------------
VERSION=""
DRY_RUN=0
for a in "$@"; do
  case "$a" in
    --dry-run) DRY_RUN=1 ;;
    v*)        VERSION="$a" ;;
    *)         echo "release: unknown arg '$a'" >&2; exit 2 ;;
  esac
done

die()  { echo "release: $*" >&2; exit 1; }
note() { echo ">> $*"; }
run()  { # execute (or, in dry-run, just print) a command
  if [[ $DRY_RUN == 1 ]]; then echo "   DRY  $*"; else echo "   RUN  $*"; eval "$*"; fi
}

[[ -n "$VERSION" ]] || die "usage: bin/release.sh vX.Y.Z [--dry-run]"
[[ "$VERSION" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]] || die "version must look like v1.2.3 (got '$VERSION')"

ROOT="$(git rev-parse --show-toplevel)"
cd "$ROOT"

# tag name for a module directory: root -> "vX.Y.Z", subdir -> "<dir>/vX.Y.Z"
tag_for_dir() { [[ "$1" == "." ]] && echo "$VERSION" || echo "$1/$VERSION"; }

# ---- preflight -----------------------------------------------------------------
[[ "$(git rev-parse --abbrev-ref HEAD)" == "main" ]] || die "must run on main (on '$(git rev-parse --abbrev-ref HEAD)')"
if ! git diff --quiet || ! git diff --cached --quiet; then
  die "working tree not clean — commit or stash first"
fi
git fetch origin --tags --quiet
[[ "$(git rev-parse main)" == "$(git rev-parse origin/main)" ]] || die "local main != origin/main — pull/push first"

# fetch brand-new tags directly from GitHub (skip proxy cache + sumdb) during go get
export GOPRIVATE="$MODPATH_ROOT"
export GOFLAGS="-mod=mod"

# ---- discover modules (parallel arrays: DIRS[i] has module path MODS[i]) --------
DIRS=(); MODS=()
while IFS= read -r gm; do
  d="$(dirname "$gm")"; d="${d#./}"
  m="$(awk 'NR==1{print $2}' "$gm")"
  DIRS+=("$d"); MODS+=("$m")
done < <(find . -name go.mod -not -path './.git/*' -not -path './.local/*' -not -path '*-claude-meta/*' | sort)
[[ ${#DIRS[@]} -gt 0 ]] || die "no go.mod modules found"

# module path -> dir (linear scan; module count is tiny)
dir_of_modpath() {
  local want="$1" i
  for i in "${!MODS[@]}"; do [[ "${MODS[$i]}" == "$want" ]] && { echo "${DIRS[$i]}"; return; }; done
}

# in-repo deps of a module dir: dirs it requires (via go.mod require lines)
deps_of_dir() {
  local d="$1" tok dep
  # module path tokens under our repo root, excluding the module's own line
  grep -oE "${MODPATH_ROOT}(/[A-Za-z0-9._/-]+)?" "$d/go.mod" 2>/dev/null \
    | sort -u | while IFS= read -r tok; do
        [[ "$tok" == "$(awk 'NR==1{print $2}' "$d/go.mod")" ]] && continue  # skip self
        dep="$(dir_of_modpath "$tok")"
        [[ -n "$dep" ]] && echo "$dep"
      done
}

# ---- refuse if any target tag already exists (never clobber / re-tag) -----------
for d in "${DIRS[@]}"; do
  t="$(tag_for_dir "$d")"
  git rev-parse -q --verify "refs/tags/$t" >/dev/null 2>&1 && die "tag $t already exists — choose a higher version"
done

# ---- refuse a version <= the current highest released version -------------------
highest="$(git tag | grep -oE 'v[0-9]+\.[0-9]+\.[0-9]+$' | sort -V | tail -1 || true)"
if [[ -n "$highest" ]]; then
  top="$(printf '%s\n%s\n' "$highest" "$VERSION" | sort -V | tail -1)"
  [[ "$top" == "$VERSION" && "$VERSION" != "$highest" ]] || die "version $VERSION is not greater than the current highest $highest"
fi
note "releasing $VERSION (current highest: ${highest:-none})  dry_run=$DRY_RUN"

# ---- topological order (Kahn): a module comes after all its in-repo deps --------
ORDER=()
emitted() { local x="$1" e; for e in ${ORDER[@]+"${ORDER[@]}"}; do [[ "$e" == "$x" ]] && return 0; done; return 1; }
remaining=("${DIRS[@]}")
while [[ ${#remaining[@]} -gt 0 ]]; do
  NEXT=(); made_progress=0
  for d in "${remaining[@]}"; do
    ready=1
    while IFS= read -r dep; do [[ -n "$dep" ]] && ! emitted "$dep" && ready=0; done < <(deps_of_dir "$d")
    if [[ $ready == 1 ]]; then ORDER+=("$d"); made_progress=1; else NEXT+=("$d"); fi
  done
  if [[ ${#NEXT[@]} -eq 0 ]]; then remaining=(); else remaining=("${NEXT[@]}"); fi
  [[ $made_progress == 0 && ${#remaining[@]} -gt 0 ]] && die "dependency cycle among modules: ${remaining[*]}"
done
note "release order: ${ORDER[*]}"

# ---- release each module in dependency order -----------------------------------
for d in "${ORDER[@]}"; do
  t="$(tag_for_dir "$d")"
  note "module '${d}'  ->  tag ${t}"

  # bump every in-repo dependency to VERSION (its tag is already pushed by now)
  changed=0
  while IFS= read -r dep; do
    [[ -z "$dep" ]] && continue
    deppath="$(awk 'NR==1{print $2}' "$dep/go.mod")"   # module path of the dep dir
    run "(cd '$d' && go get '${deppath}@${VERSION}')"
    changed=1
  done < <(deps_of_dir "$d")

  if [[ $changed == 1 ]]; then
    run "(cd '$d' && go mod tidy)"
    run "git add '$d/go.mod' '$d/go.sum'"
    run "git commit -m 'release: bump ${d} in-repo deps to ${VERSION}'"
    run "git push origin main"
  fi

  run "git tag -a '$t' -m 'release ${d} ${VERSION}'"
  run "git push origin 'refs/tags/$t'"
done

note "done. tagged: $(for d in "${ORDER[@]}"; do printf '%s ' "$(tag_for_dir "$d")"; done)"
