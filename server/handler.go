package server

import (
	"bufio"
	"net"

	"github.com/kotoyuuko/bronya/config"
)

// Handler 请求处理器
func Handler(conn net.Conn) {
	defer conn.Close()

	req := &Request{
		Scanner: bufio.NewScanner(conn),
	}

	req.Parse()

	vhost, _ := config.SearchVhost(req.Host)

	ctx := &Context{
		Vhost: vhost,
		Req:   req,
		Res:   make(chan interface{}),
		Err:   make(chan error),
	}
	go ctx.Exec()

	for {
		select {
		case res := <-ctx.Res:
			switch res.(type) {
			case error:
				DoResponse(conn, ErrorResponse(500, "Error"))
				break
			case *Response:
				DoResponse(conn, res.(*Response))
				break
			default:
				DoResponse(conn, ErrorResponse(500, "Error"))
			}
		case err := <-ctx.Err:
			DoResponse(conn, ErrorResponse(500, err.Error()))
		}
	}
}
