.PHONY: clean

clean:
	rm -rf bin
	rm -rf docs/*

build: clean
	mkdir -p bin
	go build -o bin/wm ./main.go

test: build
	go test ./...

run:
	./bin/wm
	

generate-docs: build
	mkdir -p docs
	./bin/wm docs
	

fmt:
	go fmt ./... && go vet ./...
	
all: fmt clean generate-docs run