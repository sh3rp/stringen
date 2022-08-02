all: build

clean: 
	rm -rf target
	mkdir -p target/bin

build:
	go mod tidy
	go build -o target/bin/sg cmd/sg/main.go
	go build -o target/bin/stringend cmd/stringend/main.go

install: build
	sudo cp target/bin/sg /usr/local/bin
	sudo cp target/bin/stringend /usr/local/bin

PHONY: .all .build .install