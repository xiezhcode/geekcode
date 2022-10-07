package httphandler

import (
	"net/http"
	"os"
	"strconv"

	"github.com/xiezhcode/geekcode/gohttpserver/headerhandler"
	"github.com/xiezhcode/geekcode/gohttpserver/logger"
)

func Healthz(response http.ResponseWriter, request *http.Request) {
	header := headerhandler.Header{}
	// 将 Request 中的请求转换成 Header 结构体
	header.HeaderToStruct(request.Header)
	// 将 Header 结构体的内容写入 Response Header
	header.WriteHeader(response.Header())
	// 获取系统环境变量 VERSION 并写入 Response Header 中
	header.SetHeader(response.Header(), "VERSION", os.Getenv("VERSION"))
	// 设置返回码为 200
	response.WriteHeader(http.StatusOK)
	// 返回 200 给客户端
	_, err := response.Write([]byte("200"))
	if err != nil {
		logger.Logger.Panicln(err)
		return
	}
	logger.Logger.Println("客户端 IP 地址为：" + header.GetRemoteIP(request) + ", 响应码为：" + strconv.Itoa(http.StatusOK))
}

// Stop 关闭 HTTP 服务器
func Stop(response http.ResponseWriter, request *http.Request) {
	_, err := response.Write([]byte("goodbye..."))
	if err != nil {
		logger.Logger.Panicln(err)
		return
	}
	logger.Logger.Println("http server has stopped")
	os.Exit(1)
}
