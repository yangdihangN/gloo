#!/bin/bash -ex

go run modvendor.go -copy="**/*.proto" -v

for f in $(find projects/ -name "*.proto" -or -name "solo-kit.json"); do
d=$(dirname $f)
mkdir -p  vendor/github.com/solo-io/gloo/$d
cp $f vendor/github.com/solo-io/gloo/$f
done
