.PHONY: deps build clean test

# package info
USER_DIR = src/github.com/t2y
PKG_DIR = ${USER_DIR}/issue_scanner_with_O_DIRECT

# golang info
GOPATH=$(shell pwd)
export GOPATH


all: build

deps:
	mkdir -p ${USER_DIR}
	if [ ! -L ${PKG_DIR} ]; then \
		(cd ${USER_DIR} && ln -s ${GOPATH} .) \
	fi
	go get -v -d .

build:
	go build -o main main.go

clean:
	rm -f main

test:
	go test -v -cpu 4 -race *.go
