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

for i in $(seq 0 $((n-1)))
do
    make ID=$i
done
for i in $(seq 0 $((n-1)))
do
    ./build/node-$i $i >> ./build/log &
done