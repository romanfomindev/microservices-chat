LOCAL_BIN:=$(CURDIR)/bin

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/spf13/cobra-cli@latest

lint:
	GOBIN=$(LOCAL_BIN) golangci-lint run ./... --config .golangci.pipeline.yaml

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-chat-api

generate-chat-api:
	mkdir -p pkg/chat_v1
	protoc --proto_path api/chat_v1 \
	--go_out=pkg/chat_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/chat_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/chat_v1/chat.proto


#go run cmd/main.go auth login -e user33@mail.ru -p secret
#go run cmd/main.go chat create -t eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDE0OTMxOTIsImVtYWlsIjoidXNlcjMzQG1haWwucnUiLCJyb2xlIjoiVVNFUiJ9.BCHVBcUKaFM6kknGjzkxIr0SOChz_98b7Nxd5rLZAkc -n chat1 -u user33@mail.ru -u secret
#go run cmd/main.go chat connect -t eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDE0OTMxOTIsImVtYWlsIjoidXNlcjMzQG1haWwucnUiLCJyb2xlIjoiVVNFUiJ9.BCHVBcUKaFM6kknGjzkxIr0SOChz_98b7Nxd5rLZAkc -c 5