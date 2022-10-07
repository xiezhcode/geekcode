// Package httphandler
// httphandler 包用于定义 HTTP 请求的 Handler
// httphandler.go 源文件中定义了两个方法：
//  1. Healthz：用于处理客户端的正常请求，该 Handler 函数会将 Request Header 的内容写入 Response Header，并且打印客户端的 IP 地址和响应码
//  2. Stop：用于停止 HTTP Server
package httphandler
