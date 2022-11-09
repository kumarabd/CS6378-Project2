lines=$(ps aux | grep ./build | awk '{print $2}')
for line in $lines;
do
    kill $line
done