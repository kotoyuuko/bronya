package server

import (
	"io/ioutil"
	"mime"
	"os"
	"path"

	"github.com/kotoyuuko/bronya/config"
	"github.com/kotoyuuko/bronya/logger"
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
	var files []string
	if ctx.Req.File != "/" {
		files = append(files, ctx.Req.File)
	} else {
		for _, index := range ctx.Vhost.Index {
			files = append(files, "/"+index)
		}
	}
	for _, file := range files {
		if pathExist(ctx.Vhost.Root+file) && !isDir(ctx.Vhost.Root+file) {
			fileContent, err := ioutil.ReadFile(ctx.Vhost.Root + file)
			if err != nil {
				logger.Warning.Println(err)
				ctx.Res <- ErrorResponse(500, "Error")
				return
			}

			response := &Response{
				Code:    200,
				Charset: "utf-8",
				MIME:    mime.TypeByExtension(path.Ext(ctx.Vhost.Root + file)),
				Content: string(fileContent),
			}

			if ctx.Req.Gzip {
				response.GzipEncode()
			}

			ctx.Res <- response
			break
		}
	}
	ctx.Res <- ErrorResponse(404, "Not Found")
}

func pathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func isDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
