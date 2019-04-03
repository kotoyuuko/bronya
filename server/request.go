package server

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/kotoyuuko/bronya/logger"
)

// Request 存储请求信息
type Request struct {
	ID         uint16
	Reader     *bufio.Reader
	Headers    []string
	KeepConn   bool
	Host       string
	Port       string
	Method     string
	RequestURI string
	Proto      string
	File       string
	Querys     string
	Gzip       bool
	Length     int
	Body       string
}

// ParseHeader 解析 HTTP 头部信息
func (req *Request) ParseHeader() {
	i := 0
	for {
		if line, _, err := req.Reader.ReadLine(); err != io.EOF {
			ln := string(line)

			req.Headers = append(req.Headers, ln)
			if i == 0 {
				req.Method = strings.Fields(ln)[0]
				req.RequestURI = strings.Fields(ln)[1]
				req.Proto = strings.Fields(ln)[2]
			}
			if strings.HasPrefix(ln, "Host") {
				hostWithPort := strings.Split(strings.Fields(ln)[1], ":")
				req.Host = hostWithPort[0]
				req.Port = hostWithPort[1]
			}
			if strings.HasPrefix(ln, "Accept-Encoding") {
				if strings.Index(ln, "gzip") > 0 {
					req.Gzip = true
				}
			}
			if strings.HasPrefix(ln, "Content-Length") {
				req.Length, err = strconv.Atoi(strings.Fields(ln)[1])
				if err != nil {
					req.Length = 0
				}
			}
			if strings.Contains(ln, "keep-alive") {
				req.KeepConn = true
			}
			if ln == "" {
				break
			}
			i++
		} else {
			break
		}
	}
	uri := strings.Split(req.RequestURI, "?")
	req.File = uri[0]
	if len(uri) > 1 {
		req.Querys = uri[1]
	}

	logger.Info.Println(req.Method, req.Host, req.Port, req.RequestURI)
}

// ParseBody 解析 HTTP 包内容信息
func (req *Request) ParseBody() {
	for i := 0; i < req.Length; i++ {
		b, err := req.Reader.ReadByte()
		if err != nil {
			break
		}
		req.Body += string(b)
	}
}
