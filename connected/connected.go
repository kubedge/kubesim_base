/*
Copyright 2018 Kubedge

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package connected

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const DEFAULT_SIM_NAME = "KUBESIM"
const DEFAULT_CONNECTED_UE_FILE = "/etc/kubedge/connected_ue.yaml"
const DEBUG_LOG = false

type Connecteddata struct {
	Connected map[string]string `yaml:"connected"`
}

func (conn *Connecteddata) Readconnectvalues(simname string, filename string) {
	if simname == "" {
		simname = DEFAULT_SIM_NAME
	}
	if filename == "" {
		filename = DEFAULT_CONNECTED_UE_FILE
	}

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("%s: %s: Read: %v", simname, filename, err)
	}

	err = yaml.Unmarshal(yamlFile, conn)
	if err != nil {
		log.Fatalf("%s: %s: Unmarshal: %v", simname, filename, err)
	}

	if DEBUG_LOG {
		log.Printf("%s: %s contains: connected=%s", simname, filename, conn.Connected)
	}

}
