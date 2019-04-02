package server

import (
	"net"

	"github.com/kotoyuuko/bronya/config"
	"github.com/kotoyuuko/bronya/logger"
)

// Fire 重装小兔-19C
func Fire() {
	listen(config.Config.Listen, config.Config.Port)
}

func listen(addr, port string) {
	listener, err := net.Listen("tcp", addr+":"+port)
	if err != nil {
		logger.Error.Fatalln(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error.Println(err)
		}

		go Handler(conn)
	}
}
