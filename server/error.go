package server

import "strconv"

// ErrorResponse 生成错误所需的 Response
func ErrorResponse(code int, msg string) *Response {
	response := &Response{
		Code:    code,
		Charset: "utf-8",
		MIME:    "text/html",
		Content: "<h1>Bronya Boom!</h1><h4>Code " + strconv.Itoa(code) + "</h4><p>" + msg + "</p>",
	}
	return response
}
