#!/bin/bash

# 실행 명령 확인
COMMAND=$(stat -c "%n" "$0")
CMDNAME=$(basename "${COMMAND}")
CMDDIR=$(readlink -fn "$(dirname "${COMMAND}")")

DT=$(date "+%Y%m%d_%H%M%S")

${CMDDIR}/bin/tls_client -c ${CMDDIR}/config-client.yaml -f 1 -s "smart-gw-01" -n  ${DT}.avi apd_sample4_1280x720_h264.avi

