#!/bin/bash
if [ "$1" == "" ]; then
	echo "Needs a comment"
	exit
fi
. pack_assets.sh
git add -A .
git commit -m "$1"
git push origin master
