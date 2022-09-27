#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

echo "Building docker image based off of most recent local commits of savannahnode and coreth"

SAVANODE_REMOTE="git@github.com:SavaLabs/savannahnode.git"
CORETH_REMOTE="git@github.com:SavaLabs/coreth.git"
DOCKERHUB_REPO="savannahlabs/savannahnode"

DOCKER="${DOCKER:-docker}"
SCRIPT_DIRPATH=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)
ROOT_DIRPATH="$(dirname "${SCRIPT_DIRPATH}")"

SAVA_LABS_RELATIVE_PATH="src/github.com/SavaSabs"
EXISTING_GOPATH="$GOPATH"

export GOPATH="$SCRIPT_DIRPATH/.build_image_gopath"
WORKPREFIX="$GOPATH/src/github.com/SavaLabs"

# Clone the remotes and checkout the desired branch/commits
SAVANODE_CLONE="$WORKPREFIX/savannahnode"
CORETH_CLONE="$WORKPREFIX/coreth"

# Replace the WORKPREFIX directory
rm -rf "$WORKPREFIX"
mkdir -p "$WORKPREFIX"


SAVANODE_COMMIT_HASH="$(git -C "$EXISTING_GOPATH/$SAVA_LABS_RELATIVE_PATH/savannahnode" rev-parse --short HEAD)"
CORETH_COMMIT_HASH="$(git -C "$EXISTING_GOPATH/$SAVA_LABS_RELATIVE_PATH/coreth" rev-parse --short HEAD)"

git config --global credential.helper cache

git clone "$SAVANODE_REMOTE" "$SAVANODE_CLONE"
git -C "$SAVANODE_CLONE" checkout "$SAVANODE_COMMIT_HASH"

git clone "$CORETH_REMOTE" "$CORETH_CLONE"
git -C "$CORETH_CLONE" checkout "$CORETH_COMMIT_HASH"

CONCATENATED_HASHES="$SAVANODE_COMMIT_HASH-$CORETH_COMMIT_HASH"

"$DOCKER" build -t "$DOCKERHUB_REPO:$CONCATENATED_HASHES" "$WORKPREFIX" -f "$SCRIPT_DIRPATH/local.Dockerfile"
