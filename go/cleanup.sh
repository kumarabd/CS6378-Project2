lines=$(ps aux | grep ./build | awk '{print $2}')
#rm -rf build/*
for line in $lines;
do
    kill $line
done