all: build

clean: 
	rm -rf target
	mkdir -p target/bin

build:
	go mod tidy
	go build -o target/bin/stringen cmd/sg/main.go
	go build -o target/bin/stringend cmd/stringend/main.go

install: build
	cp target/bin/stringen $(HOME)/bin
	cp target/bin/stringend $(HOME)/bin

PHONY: .all .build .install
