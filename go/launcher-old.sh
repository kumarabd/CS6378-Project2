counter=0
n=0
while IFS= read -r line; do
    if [[ $line != \#* ]] && ! [[ -z "$line" ]]
    then
        (( counter++ ))
        case $counter in
        1)
            for num in $line
            do 
                n=$num
                break
            done;;
        *) continue;;
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