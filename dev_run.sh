#!/bin/bash
if [ "$1" == "first" ]; then
	go get -u
	. pack_assets.sh
	cp settings.example.cfg settings.cfg
	go run main.go /api/install
else
	go get -u
	. pack_assets.sh
	go run main.go
fi
