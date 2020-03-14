
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/uc/uc.proto

.PHONY: build
build: proto

	GOOS=linux
	GOARCH=amd64
	go build -o uc-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t uc-srv:latest
