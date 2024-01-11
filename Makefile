.PHONY: all build test clean 

all: build

build:
	go build
	./learn1

test: 
	go test -v

clean: 
	rm -f ./learn1


