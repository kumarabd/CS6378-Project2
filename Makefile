ID?=0
all:
	g++ -pthread -std=c++11 main.cpp node.cpp output.cpp mutex.cpp application.cpp channel.cpp -o build/node-$(ID)

test:
	g++ -pthread -std=c++11 main.cpp node.cpp output.cpp mutex.cpp application.cpp channel.cpp -o build/node-test
	./build/node-test 0
	
clean:
	rm -f build/*