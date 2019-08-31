all: clean test build

test:
	go test -v ./...

clean:
	rm -rf bin/main

build: clean
	go build -o bin/main cmd/main.go

run: clean
	./bin/main
