package server

import (
	"bufio"
	"strings"
)

// Request 存储请求信息
type Request struct {
	ID         uint16
	Scanner    *bufio.Scanner
	Header     []string
	KeepConn   bool
	Host       string
	Port       string
	Method     string
	RequestURI string
	Proto      string
}

// Parse 解析 HTTP 头部信息
func (req *Request) Parse() {
	i := 0
	for req.Scanner.Scan() {
		ln := req.Scanner.Text()
		req.Header = append(req.Header, ln)
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
		if strings.Contains(ln, "keep-alive") {
			req.KeepConn = true
		}
		if ln == "" {
			break
		}
		i++
	}
}
