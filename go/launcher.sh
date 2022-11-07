#!/bin/bash

# Change this to your netid
netid=axk200020
passwd="buvkiz-dihrIv-3sezhy"

# Root directory of your project
PROJDIR=/home/012/a/ax/$netid/aos-project-2

# Directory where the config file is located on your local system
CONFIGLOCAL=/Users/abishekk/Documents/github-projects/CS6378-Project2/go/config.txt

# Directory your java classes are in
BINDIR=/home/012/a/ax/$netid/aos-project-2/build

n=0

cat $CONFIGLOCAL | sed -e "s/#.*//" | sed -e "/^\s*$/d" |
(
    read i
    i=$(echo $i | awk '{print $1;}')
    while [[ $n -lt $i ]]
    do
    	read line
    	p=$( echo $line | awk '{ print $1 }' )
        host=$( echo $line | awk '{ print $2 }' )

        osascript -e 'tell app "Terminal"
do script "ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no '$netid'@'$host' go version; exec bash"
end tell' &

	    #do script "ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no '$netid'@$'host' -l cd /home/012/a/ax/axk200020/aos-project-2; make ID='$i'; ./build/node-'$i' '$i' ;exec bash"

        n=$(( n + 1 ))
    done
)
