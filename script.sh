#!/usr/bin/env bash
set -e pipefail

WDIR=$(pwd)

function build_executable() {
    cd src
    go build -o ../bin/k8s-filegen
    echo "Created Executable"
}

function move_to_path() {
    k8s-gen || cd "${WDIR}" && make build
    cd "${WDIR}/bin"
    mv k8s-filegen /usr/local/bin
    echo "Moved Executable"
}

function remove() {
    rm -rf /usr/local/bin/k8s-filegen
    echo "Deleted Executable"
}
