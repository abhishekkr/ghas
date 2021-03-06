#!/usr/bin/env bash

set -ex

PROJECT_DIR=$(dirname $0)
BUILD_DIR="${PROJECT_DIR}/out"
BIN_NAME="ghas"

buildBinaries(){
  [[ ! -d "${BUILD_DIR}" ]] && mkdir "${BUILD_DIR}"

  go build -o "${BUILD_DIR}/${BIN_NAME}-linux-amd64" ghas.go

  GOOS=windows go build -o "${BUILD_DIR}/${BIN_NAME}-windows-amd64" ghas.go

  GOOS=darwin go build -o "${BUILD_DIR}/${BIN_NAME}-darwin-amd64" ghas.go
}

##### main
buildBinaries
