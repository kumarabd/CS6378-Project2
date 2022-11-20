#!/bin/bash

# Change this to your netid
netid=axk200020

# Root directory of your project
PROJDIR=/home/012/a/ax/$netid/aos-project-2
PROJDIR_LOCAL=/Users/abishekk/Documents/github-projects/CS6378-Project2/go

# Directory where the config file is located on your local system
CONFIGLOCAL=/Users/abishekk/Downloads/CS6378-Project2/go/config.txt

# Directory your java classes are in
BINDIR=/home/012/a/ax/$netid/aos-project-2/build

n=0
unameOut="$(uname -s)"
cat $CONFIGLOCAL | sed -e "s/#.*//" | sed -e "/^\s*$/d" |
(
    read i
    i=$(echo $i | awk '{print $1;}')
    echo $i
    while [[ $n -lt $i ]]
    do
    	read line
    	p=$( echo $line | awk '{ print $1 }' )
        host=$( echo $line | awk '{ print $2 }' )
        case "${unameOut}" in
            Linux*)     gnome-terminal --tab -- 'ssh -i ~/.ssh/utd_rsa -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -Y '$netid'@'$host' \"cd '$PROJDIR'; make ID='$n';./build/node-'$n'\"'&;;
            Darwin*)    osascript -e 'tell app "Terminal"
do script "ssh -i ~/.ssh/utd_rsa -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -Y '$netid'@'$host' \"cd '$PROJDIR'; make ID='$n';./build/node-'$n'\""
end tell' &;;
        esac
    fi
done < config.txt

for i in $(seq 1 $n)
do
    make ID=$i
done
for i in $(seq 1 $n)
do
    ./build/node-$i $i &
done