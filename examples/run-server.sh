#!/bin/bash

# 실행 명령 확인
COMMAND=$(stat -c "%n" "$0")
CMDNAME=$(basename "${COMMAND}")
CMDDIR=$(readlink -fn "$(dirname "${COMMAND}")")

${CMDDIR}/bin/tls_server -c ${CMDDIR}/config-server.yaml
