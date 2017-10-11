#!/bin/bash

WDIR=$(pwd)

mkdir -p build/webapp
cd build

env GOARCH=amd64 GOOS=windows go build -ldflags "-s -w" ../cli

cd $WDIR/webapp
npm install
ng build -prod --aot --output-path $WDIR/build/webapp/dist

cd $WDIR
tar czvf windows.tar.gz build/

rm -rf build
