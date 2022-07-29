#!/bin/sh

# delete old binary
sudo rm /usr/local/bin/locvis

# build
go build

# move
sudo mv ./locvis /usr/local/bin