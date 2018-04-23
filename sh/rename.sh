#
# Rename Shell Script
# TODO: rename `gweb` to other name so easy to use 
# 

#!/usr/bin

# make sure you do this script in folder `path/to/gewb/sh`

NewName=$1

if [ $NewName = '' ]; then
	echo "Need a arg as NewName for your Project"
	exit
fi

PATH=`cd ../ && pwd`
echo "Command Running: $PATH"

sed -i "s/gewb/$NewName/g" `grep gewb -rl $PATH`