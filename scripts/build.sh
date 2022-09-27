#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# Savannahnode root folder
SAVANODE_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )
# Load the versions
source "$SAVANODE_PATH"/scripts/versions.sh
# Load the constants
source "$SAVANODE_PATH"/scripts/constants.sh

# Download dependencies
echo "Downloading dependencies..."
go mod download

# Build savannahnode
"$SAVANODE_PATH"/scripts/build_savannahnode.sh

# Build coreth
"$SAVANODE_PATH"/scripts/build_coreth.sh

# Exit build successfully if the binaries are created
if [[ -f "$savanode_path" && -f "$evm_path" ]]; then
        echo "Build Successful"
        exit 0
else
        echo "Build failure" >&2
        exit 1
fi
