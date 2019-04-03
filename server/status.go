package server

// HTTPStatusCode 存储各种 HTTP 状态码描述
var HTTPStatusCode map[int]string

func init() {
	HTTPStatusCode = make(map[int]string)

	// 1xx - Informational
	HTTPStatusCode[100] = "Continue"
	HTTPStatusCode[101] = "Switching Protocols"
	HTTPStatusCode[102] = "Processing"

	// 2xx - Success
	HTTPStatusCode[200] = "OK"
	HTTPStatusCode[201] = "Created"
	HTTPStatusCode[202] = "Accepted"
	HTTPStatusCode[203] = "Non-authoritative Information"
	HTTPStatusCode[204] = "No Content"
	HTTPStatusCode[205] = "Reset Content"
	HTTPStatusCode[206] = "Partial Content"
	HTTPStatusCode[207] = "Multi-Status"
	HTTPStatusCode[208] = "Already Reported"
	HTTPStatusCode[226] = "IM Used"

	// 3xx - Redirection
	HTTPStatusCode[300] = "Multiple Choices"
	HTTPStatusCode[301] = "Moved Permanently"
	HTTPStatusCode[302] = "Found"
	HTTPStatusCode[303] = "See Other"
	HTTPStatusCode[304] = "Not Modified"
	HTTPStatusCode[305] = "Use Proxy"
	HTTPStatusCode[307] = "Temporary Redirect"
	HTTPStatusCode[308] = "Permanent Redirect"

	// 4xx - Client Error
	HTTPStatusCode[400] = "Bad Request"
	HTTPStatusCode[401] = "Unauthorized"
	HTTPStatusCode[402] = "Payment Required"
	HTTPStatusCode[403] = "Forbidden"
	HTTPStatusCode[404] = "Not Found"
	HTTPStatusCode[405] = "Method Not Allowed"
	HTTPStatusCode[406] = "Not Acceptable"
	HTTPStatusCode[407] = "Proxy Authentication Required"
	HTTPStatusCode[408] = "Request Timeout"
	HTTPStatusCode[409] = "Conflict"
	HTTPStatusCode[410] = "Gone"
	HTTPStatusCode[411] = "Length Required"
	HTTPStatusCode[412] = "Precondition Failed"
	HTTPStatusCode[413] = "Payload Too Large"
	HTTPStatusCode[414] = "Request-URI Too Long"
	HTTPStatusCode[415] = "Unsupported Media Type"
	HTTPStatusCode[416] = "Requested Range Not Satisfiable"
	HTTPStatusCode[417] = "Expectation Failed"
	HTTPStatusCode[418] = "I'm a teapot"
	HTTPStatusCode[421] = "Misdirected Request"
	HTTPStatusCode[422] = "Unprocessable Entity"
	HTTPStatusCode[423] = "Locked"
	HTTPStatusCode[424] = "Failed Dependency"
	HTTPStatusCode[426] = "Upgrade Required"
	HTTPStatusCode[428] = "Precondition Required"
	HTTPStatusCode[429] = "Too Many Requests"
	HTTPStatusCode[431] = "Request Header Fields Too Large"
	HTTPStatusCode[444] = "Connection Closed Without Response"
	HTTPStatusCode[451] = "Unavailable For Legal Reasons"
	HTTPStatusCode[499] = "Client Closed Request"

	// 5xx - Server Error
	HTTPStatusCode[500] = "Internal Server Error"
	HTTPStatusCode[501] = "Not Implemented"
	HTTPStatusCode[502] = "Bad Gateway"
	HTTPStatusCode[503] = "Service Unavailable"
	HTTPStatusCode[504] = "Gateway Timeout"
	HTTPStatusCode[505] = "HTTP Version Not Supported"
	HTTPStatusCode[506] = "Variant Also Negotiates"
	HTTPStatusCode[507] = "Insufficient Storage"
	HTTPStatusCode[508] = "Loop Detected"
	HTTPStatusCode[510] = "Not Extended"
	HTTPStatusCode[511] = "Network Authentication Required"
	HTTPStatusCode[599] = "Network Connect Timeout Error"
}
