
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/dotc/dotc.proto

.PHONY: build
build: proto

	GOOS=linux
	GOARCH=amd64
	go build -o dotc-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t dotc-srv:latest
