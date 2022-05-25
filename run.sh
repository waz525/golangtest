#!/bin/sh

golang="$1"

if [ "$golang" == "" ]
then
	echo "Usage: $0 FileName"
else
	echo "Run ${golang} ..."
	echo '--------------------------'
	shift
	go run ${golang} $*
fi
