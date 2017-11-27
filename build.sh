#!/bin/bash
## VFT Build script v1, written by pry0cc 27/11/2017

startdir=$PWD

mkdir -p ./bin/

if [ -z ${GOPATH+x} ]
then
	echo "No GOPATH set, creating one"
	mkdir -p $HOME/tmp/golang/
	export GOPATH="$HOME/tmp/golang/"
else
	echo "GOPATH set, continuing."
fi

cd $GOPATH
mkdir -p ./src/github.com/bbriggs
echo "Cloning vft"
git clone git@gitlab.s-3.tech:fraq/vft.git ./src/github.com/bbriggs/vft
cd $GOPATH/src/github.com/bbriggs/vft/cmd/vft
echo "Installing dependencies"
go get ./...
echo "Building vft"
go build -o $startdir/bin/vft ./vft.go
echo "vft compiled to $startdir/bin/vft, enjoy!"
