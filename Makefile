ID?=0
all:
	g++ -pthread -std=c++11 main.cpp node.cpp output.cpp mutex.cpp application.cpp channel.cpp -o build/app-$(ID)

test:
	g++ -pthread -std=c++11 main.cpp node.cpp output.cpp mutex.cpp application.cpp channel.cpp -o build/app-test
	./build/app-test 0
	
clean:
	rm -f build/*