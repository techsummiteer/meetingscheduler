.PHONY: all build test clean 

all: build

build:
	go build
	./meetingscheduler

test: 
	go test -v

clean: 
	rm -f ./learn1


