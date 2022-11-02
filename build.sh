#!/usr/bin/bash

archs=(amd64 arm64 ppc64le ppc64 s390x)

for arch in ${archs[@]}
do
        env GOOS=linux GOARCH=${arch} go build -o binaries/sandpiles-${arch}
done

env GOOS=windows GOARCH=amd64 go build -o binaries/sandpiles-amd64.exe

env GOOS=darwin GOARCH=amd64 go build -o binaries/sandpiles-amd64-darwin



