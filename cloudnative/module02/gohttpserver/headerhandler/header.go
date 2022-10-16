package headerhandler

import (
	"net/http"
	"reflect"
	"strings"
)

type Header struct {
	Accept                  string `field:"Accept"`
	AcceptEncoding          string `field:"Accept-Encoding"`
	AcceptLanguage          string `field:"Accept-Language"`
	CacheControl            string `field:"Cache-Control"`
	Connection              string `field:"Connection"`
	SecChUa                 string `field:"Sec-Ch-Ua"`
	SecChUaMobile           string `field:"Sec-Ch-Ua-Mobile"`
	SecChUaPlatform         string `field:"Sec-Ch-Ua-Platform"`
	SecFetchDest            string `field:"Sec-Fetch-Dest"`
	SecFetchMode            string `field:"Sec-Fetch-Mode"`
	SecFetchSite            string `field:"Sec-Fetch-Site"`
	SecFetchUser            string `field:"Sec-Fetch-User"`
	UpgradeInsecureRequests string `field:"Upgrade-Insecure-Requests"`
	UserAgent               string `field:"User-Agent"`
	XForwardFor             string `field:"X-Forward-For"`
	XRealIP                 string `field:"X-Real-IP"`
}

// HeaderToStruct 获取 HTTP Header 的内容转换成 Header 结构体
// 为了避免后续 Header 结构体新增字段导致要修改此方法，所以使用反射获取结构体的字段进行赋值
func (h *Header) HeaderToStruct(header http.Header) {
	hType := reflect.TypeOf(h)
	hValue := reflect.ValueOf(h)

	// 注意：下方调用各类方法时，要使用非指针类型，但由于方法的接收者是指针，所以需要调用 Elem() 方法将指针类型转换成非指针类型
	for i := 0; i < hType.Elem().NumField(); i++ {
		hValue.Elem().Field(i).SetString(header.Get(hType.Elem().Field(i).Tag.Get("field")))
	}
}

// WriteHeader 将 Header 结构体的内容写入 Response Header
// 为了避免后续 Header 结构体新增字段导致要修改此方法，所以使用反射获取结构体字段的标签、值进行赋值
func (h *Header) WriteHeader(header http.Header) {
	hType := reflect.TypeOf(h)
	hValue := reflect.ValueOf(h)

	for i := 0; i < hType.Elem().NumField(); i++ {
		tag := hType.Elem().Field(i).Tag.Get("field")
		value := hValue.Elem().Field(i).String()
		if value != "" {
			header.Set(tag, value)
		}
	}
}

// SetHeader 为 Response Header 设置新的字段
func (h *Header) SetHeader(header http.Header, key, value string) {
	header.Set(key, value)
}

// GetRemoteIP 获取客户端 IP 地址
func (h *Header) GetRemoteIP(request *http.Request) string {
	if h.XForwardFor != "" {
		xForwardForArr := strings.Split(h.XForwardFor, ",")
		for _, element := range xForwardForArr {
			if strings.TrimSpace(element) != "" && strings.TrimSpace(element) != "unknown" {
				return strings.TrimSpace(element)
			}
		}
	}

	if h.XRealIP != "" {
		return strings.TrimSpace(h.XRealIP)
	}

	remoteAddr := strings.TrimSpace(request.RemoteAddr)

	if remoteAddr != "" {
		index := strings.LastIndex(remoteAddr, ":")
		return remoteAddr[:index]
	}

	return ""
}

// ToString 将 Struct 转换成字符串，用于输出展示
// 为了避免后续 Header 结构体新增字段导致要修改此方法，所以使用反射获取结构体字段的值
func (h *Header) ToString() string {
	hType := reflect.TypeOf(h)
	hValue := reflect.ValueOf(h)
	stringBuilder := strings.Builder{}
	stringBuilder.WriteString("{")

	for i := 0; i < hType.Elem().NumField(); i++ {
		stringBuilder.WriteString("\"" + hType.Elem().Field(i).Name + "\": ")
		if i == hType.Elem().NumField()-1 {
			stringBuilder.WriteString("[ " + hValue.Elem().Field(i).String() + " ]" + " }")
			break
		}
		stringBuilder.WriteString("[ " + hValue.Elem().Field(i).String() + " ]" + ", ")
	}

	return stringBuilder.String()
}
