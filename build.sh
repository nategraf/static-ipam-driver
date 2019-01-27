#!/usr/bin/env bash
set -e

cd "$( dirname "${BASH_SOURCE[0]}" )"
mkdir -p ./bin

for arch in "amd64" "386"; do
    for os in "linux" "windows"; do
        export GOARCH="$arch"
        export GOOS="$os"
        go build -o bin/static-ipam-driver.$os.$arch
    done
done
