package server

import (
	"github.com/kotoyuuko/bronya/config"
)

// Context 负责 channel 间通信
type Context struct {
	Vhost *config.Vhost
	Req   *Request
	Res   chan interface{}
	Err   chan error
}

// Exec 处理请求
func (ctx *Context) Exec() {
	response := &Response{
		Code:    200,
		Charset: "utf-8",
		MIME:    "text/html",
		Content: "测试",
	}
	ctx.Res <- response
}
