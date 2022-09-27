#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# Directory above this script
SAVANODE_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )
# Load the versions
source "$SAVANODE_PATH"/scripts/versions.sh
# Load the constants
source "$SAVANODE_PATH"/scripts/constants.sh

go test -race -timeout="120s" -coverprofile="coverage.out" -covermode="atomic" $(go list ./... | grep -v /mocks | grep -v proto | grep -v tests)
