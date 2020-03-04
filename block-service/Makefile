
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/block/block.proto

.PHONY: build
build: proto

	GOOS=linux
	GOARCH=amd64
	go build -o block-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t block-srv:latest
