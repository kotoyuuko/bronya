package main

import (
	"github.com/kotoyuuko/bronya/logger"
	"github.com/kotoyuuko/bronya/server"
)

func main() {
	logger.Info.Println("Bronya Starting...")
	server.Fire()
}
