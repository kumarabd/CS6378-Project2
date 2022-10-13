all:
	g++ -pthread -std=c++11 main.cpp network.cpp node.cpp output.cpp MEService.cpp -o build/app
	./build/app

clean:
	rm -f build/*