FROM alpine:3.15

ENV OPERATOR=/usr/local/bin/healthsrv \
    USER_UID=1001 \
    USER_NAME=healthsrv

# install operator binary
COPY build/_output/bin/healthsrv ${OPERATOR}

COPY build/bin /usr/local/bin

RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
