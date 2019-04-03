package server

import "strconv"

// ErrorResponse 生成错误所需的 Response
func ErrorResponse(code int, msg string) *Response {
	response := &Response{
		Code:    code,
		Content: "<h1>Bronya Boom!</h1><h4>Code " + strconv.Itoa(code) + "</h4><p>" + msg + "</p>",
	}
	response.Header("Content-Type: text/html; charset=utf-8")
	return response
}
