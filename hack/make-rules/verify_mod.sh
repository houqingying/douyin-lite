#!/usr/bin/env bash

set -e
set -u

BPFLET_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd -P)"

echo ${BPFLET_ROOT}

go mod tidy

if [ ! -z "$(git diff ${BPFLET_ROOT}/go.mod ${BPFLET_ROOT}/go.sum)" ]; then
    echo "Verify go mod error, please ensure go.mod has been tidied"
    exit -1
fi 
