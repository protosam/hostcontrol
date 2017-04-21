#!/bin/bash

if which mysqldump > /dev/null 2>&1; then
	MYSQLCHECK=$(mysqldump --no-data --compact --skip-comments hostcontrol | grep -v '^\/\*![0-9]\{5\}.*\/;$' | wc -l)
	if [ "$MYSQLCHECK" != "0" ]; then
	    echo Making an up-to-date MySQL dump
	    mysqldump --no-data hostcontrol > common/hostcontrol.sql
	fi
fi

# jteeuwen hasn't updated go-bindata since golang moved vet out of experimental
# seanmcgary made some maintenance changes

if ! which $GOPATH/bin/go-bindata; then
	echo Installing go-bindata;
	go get -u github.com/seanmcgary/go-bindata
	cd $GOPATH/src/github.com/seanmcgary/go-bindata
	pwd
	# try make...
	if ! make; then
		# if make fails, try manual build/install
		go get
		go build
		go install
	fi
	cd -
fi

if ! which $GOPATH/bin/go-bindata; then
	echo 'Dying... because I could not install go-bindata'
fi

echo Creating assets.go files...

find . -type d | grep assets | while read assetdir; do
	parentdir=$(echo $assetdir | awk -F'/' '{print $(NF-1)}')
	cd $assetdir/..
	echo $GOPATH/bin/go-bindata -pkg $parentdir -o assets.go assets/
	$GOPATH/bin/go-bindata -pkg $parentdir -o assets.go assets/
	cd -
done
