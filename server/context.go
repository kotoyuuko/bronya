package server

import (
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/kotoyuuko/bronya/config"
	"github.com/kotoyuuko/bronya/fcgi"
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
			response := &Response{
				Code: 200,
			}

			if strings.HasSuffix(file, ".php") {
				env := make(map[string]string)
				env["SCRIPT_FILENAME"] = ctx.Vhost.Root + file
				env["SERVER_SOFTWARE"] = "Bronya/1.0.0"
				env["REMOTE_ADDR"] = "127.0.0.1"
				env["QUERY_STRING"] = ctx.Req.Querys

				fcgi, err := fcgi.Dial(ctx.Vhost.Fastcgi.Network, ctx.Vhost.Fastcgi.Address)
				if err != nil {
					logger.Error.Println(err)
					ctx.Res <- ErrorResponse(502, "Bad Gateway")
					break
				}

				var resp *http.Response

				if ctx.Req.Method == "POST" {
					querys, err := url.ParseQuery(ctx.Req.Body)
					if err != nil {
						logger.Error.Println(err)
						ctx.Res <- ErrorResponse(500, "Internal Server Error")
						break
					}

					resp, err = fcgi.PostForm(env, querys)
					if err != nil {
						logger.Error.Println(err)
						ctx.Res <- ErrorResponse(502, "Bad Gateway")
						break
					}
				} else {
					resp, err = fcgi.Get(env)
					if err != nil {
						logger.Error.Println(err)
						ctx.Res <- ErrorResponse(502, "Bad Gateway")
						break
					}
				}

				for k, val := range resp.Header {
					for _, v := range val {
						response.Header(k + ": " + v)
					}
				}

				content, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					logger.Error.Println(err)
					ctx.Res <- ErrorResponse(502, "Bad Gateway")
					break
				}

				response.Content = string(content)
			} else {
				fileContent, err := ioutil.ReadFile(ctx.Vhost.Root + file)
				if err != nil {
					logger.Warning.Println(err)
					ctx.Res <- ErrorResponse(500, "Internal Server Error")
					break
				}

				response.Header("Content-Type: " + mime.TypeByExtension(path.Ext(ctx.Vhost.Root+file)))
				response.Content = string(fileContent)
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
