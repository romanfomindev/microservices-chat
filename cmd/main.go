package main

import (
	"github.com/romanfomindev/microservices-chat/cmd/root"
	"github.com/romanfomindev/microservices-chat/internal/storage"
)

func main() {
	storage.Init()
	root.Execute()
}
