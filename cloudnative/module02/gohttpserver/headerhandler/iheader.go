package headerhandler

import "net/http"

// IHeader 处理 HTTP Header 的接口
type IHeader interface {
	// HeaderToStruct 将获取到的 HTTP Header 转换成 Struct 对象
	HeaderToStruct(header http.Header)
	// WriteHeader 将 Header 结构体的内容写入 Response Header
	WriteHeader(header http.Header)
	// SetHeader 为 Response Header 设置新的字段
	SetHeader(header http.Header, key, value string)
	// GetRemoteIP 获取客户端 IP 地址
	GetRemoteIP(request *http.Request) string
	// ToString 将 Struct 转换成字符串，用于输出展示
	ToString() string
}
