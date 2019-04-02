package server

import (
	"net"
	"strconv"
)

// Response 存储响应信息
type Response struct {
	Code    int
	Charset string
	MIME    string
	Headers []string
	Content string
}

// Bytes 将响应内容转换为 Bytes 数组
func (resp *Response) Bytes() []byte {
	return []byte(resp.Content)
}

// Length 计算响应内容长度
func (resp *Response) Length() int {
	return len(resp.Content)
}

// Header 向响应数据包内添加自定义 Header
func (resp *Response) Header(header string) {
	resp.Headers = append(resp.Headers, header)
}

// DoResponse 发送响应
func DoResponse(conn net.Conn, resp *Response) {
	respPkg := "HTTP/1.1 " + strconv.Itoa(resp.Code) + "\r\n"
	respPkg += "Content-Type: " + resp.MIME + "; charset=" + resp.Charset + "\r\n"
	respPkg += "Content-Length: " + strconv.Itoa(resp.Length()) + "\r\n"

	for _, header := range resp.Headers {
		respPkg += header
	}

	respPkg += "\r\n"
	respPkg += resp.Content

	conn.Write([]byte(respPkg))
}
