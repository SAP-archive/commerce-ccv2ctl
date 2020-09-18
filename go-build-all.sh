#!/bin/bash
# Based on: https://gist.github.com/eduncan911/68775dba9d3c028181e4
#

#
PLATFORMS="darwin/amd64" # amd64 only as of go1.5
PLATFORMS="$PLATFORMS windows/amd64" # arm compilation not available for Windows
# PLATFORMS="$PLATFORMS windows/386"
PLATFORMS="$PLATFORMS linux/amd64"
# PLATFORMS="$PLATFORMS linux/386"
# PLATFORMS="$PLATFORMS linux/ppc64 linux/ppc64le"
# PLATFORMS="$PLATFORMS linux/mips64 linux/mips64le" # experimental in go1.6
# PLATFORMS="$PLATFORMS freebsd/amd64"
# PLATFORMS="$PLATFORMS netbsd/amd64" # amd64 only as of go1.6
# PLATFORMS="$PLATFORMS openbsd/amd64" # amd64 only as of go1.6
# PLATFORMS="$PLATFORMS dragonfly/amd64" # amd64 only as of go1.5
# PLATFORMS="$PLATFORMS plan9/amd64 plan9/386" # as of go1.4
# PLATFORMS="$PLATFORMS solaris/amd64" # as of go1.3


##############################################################
# Shouldn't really need to modify anything below this line.  #
##############################################################

type setopt >/dev/null 2>&1

SCRIPT_NAME=$(basename "$0")
FAILURES=""
OUTPUT="ccv2ctl"

for PLATFORM in $PLATFORMS; do
  GOOS=${PLATFORM%/*}
  GOARCH=${PLATFORM#*/}
  BIN_FILENAME="${OUTPUT}-${GOOS}-${GOARCH}"
  if [[ "${GOOS}" == "windows" ]]; then BIN_FILENAME="${BIN_FILENAME}.exe"; fi
  CMD="GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${BIN_FILENAME}"
  echo "${CMD}"
  eval "$CMD" || FAILURES="${FAILURES} ${PLATFORM}"
done

# eval errors
if [[ "${FAILURES}" != "" ]]; then
  echo ""
  echo "${SCRIPT_NAME} failed on: ${FAILURES}"
  exit 1
fi
