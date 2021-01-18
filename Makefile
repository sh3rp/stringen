all: build

clean: 
	rm -rf target
	mkdir target

build:
	go build -o target/sg cmd/sg/main.go
	go build -o target/stringend cmd/stringend/main.go

install: build
	sudo cp target/sg /usr/local/bin
	sudo cp target/stringend /usr/local/bin

PHONY: .all .build .install