#!/bin/bash

# Change this to your netid
netid=axk200020

# Root directory of your project
PROJDIR=/home/012/a/ax/$netid/aos-project-2
PROJDIR_LOCAL=/Users/abishekk/Documents/github-projects/CS6378-Project2/go

# Directory where the config file is located on your local system
CONFIGLOCAL=/Users/abishekk/Documents/github-projects/CS6378-Project2/go/config.txt

# Directory your java classes are in
BINDIR=/home/012/a/ax/$netid/aos-project-2/build

n=0
unameOut="$(uname -s)"
cat $CONFIGLOCAL | sed -e "s/#.*//" | sed -e "/^\s*$/d" |
(
    read i
    i=$(echo $i | awk '{print $1;}')
    while [[ $n -lt $i ]]
    do
    	read line
    	p=$( echo $line | awk '{ print $1 }' )
        host=$( echo $line | awk '{ print $2 }' )
        case "${unameOut}" in
            Linux*)     gnome-terminal -e 'ssh -i ~/.ssh/utd_rsa -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no '$netid'@'$host' \"cd '$PROJDIR'; make ID='$n';./build/node-'$n'\" | tee output-'$n'.txt' &;;
            Darwin*)    osascript -e 'tell app "Terminal"
do script "cd '$PROJDIR_LOCAL'; make ID='$n';./build/node-'$n' | tee ./build/output-'$n'.txt"
end tell' &;;
        esac
        n=$(( n + 1 ))
    done
)
