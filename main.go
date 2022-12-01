package main

import (
	"github.com/lohmander/blocks/core"
	"github.com/lohmander/blocks/http_"
	"github.com/lohmander/blocks/naive"
)

func main() {
	server := core.NewServer(naive.NewStore(), &http_.HTTPTransport{})
	server.Serve()
}
