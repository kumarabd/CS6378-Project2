ID?=0
all:
	go build -ldflags="-X 'main.ID=${ID}'" -o ./build/node-${ID} main.go

test:
	go build -ldflags="-X 'main.ID=${ID}'" -o ./build/node-${ID} main.go
	./node-0
	
clean:
	rm -f build/*