lines=$(ps aux | grep ./build | awk '{print $2}')
#rm -rf build/*
for line in $lines;
do
    kill $line
done

##!/bin/bash

## Change this to your netid
#netid=axk200020

## Root directory of your project
#PROJDIR=/home/012/a/ax/$netid/aos-project-2

## Directory where the config file is located on your local system
#CONFIGLOCAL=/Users/abishekk/Documents/github-projects/CS6378-Project2/go/config.txt

#n=0

#cat $CONFIGLOCAL | sed -e "s/#.*//" | sed -e "/^\s*$/d" |
#(
#    read i
#    i=$(echo $i | awk '{print $1;}')
#    while [[ $n -lt $i ]]
#    do
#    	read line
#        host=$( echo $line | awk '{ print $2 }' )

#        osascript -e 'tell app "Terminal"
#        do script "ssh -i ~/.ssh/utd_rsa -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no '$netid'@'$host' \"cd '$PROJDIR'; rm -rf ./build/*\""
#        end tell' &

#        n=$(( n + 1 ))
#    done
   
#)


#echo "Cleanup complete"
