# ---------------------------------------------------------------------------
# Makefile for GO utilities
# ---------------------------------------------------------------------------

PROJECT_DIR=$(notdir $(shell pwd))

BUILD_TAG=`git describe --tags 2>/dev/null`
LDFLAGS=-ldflags "all=-X main.version=${BUILD_TAG} -s -w"

all: get build

init:
	sed -i "s/PROJECT_NAME/${PROJECT_DIR}/g" CHANGELOG.md README.md main.go

build:
	go build ${LDFLAGS}

get:
	go get

test: fmt vet
	go test -v -cover

cover:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

fmt:
	go fmt

vet:
	go vet -v

install:
	go install ${LDFLAGS} ./...

dist: clean build
	upx -9 ${PROJECT_DIR}.exe

clean:
	go clean
