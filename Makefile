.PHONY: clean

clean:
	rm -rf bin

build: clean
	mkdir -p bin
	go build -o bin/wm ./main.go

test: build
	go test ./...

run:
	./bin/wm
	

fmt:
	go fmt ./... && go vet ./...
	
all: fmt clean build run