#!/bin/sh

OSES='linux windows darwin'
ARCHES='amd64 386'
FLAGS='-ldflags="-s -w"' # strip symbol table and disable DWARF generation

for os in $OSES; do
	for arch in $ARCHES; do
		cmd="GOOS=$os GOARCH=$arch go build -o arxor_${os}_${arch} $FLAGS"
		echo "$cmd"
		eval "$cmd"
	done
done
