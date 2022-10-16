// Package headerhandler
// headerhandler 包用于处理 Header
// iheader.go 源文件中定义了一个名为 IHeader 的接口，接口中有以下方法：
//  1. HeaderToStruct：该方法传入 Request Header，将 Header 的内容转换成 Header 结构体
//  2. WriteHeader：该方法传入 Response Header，将 Header 结构体的内容写入 Response Header
//  3. SetHeader：该方法传入三个参数，分别为：Response Header、Header 字段的名（key）和值（value），将指定的字段和其值传入 Response Header
//  4. GetRemoteIP：该方法用于获取客户端 IP 地址
//  5. ToString：该方法将 Header 结构体的内容转换成指定格式的字符串，用于输出展示
//
// header.go 源文件中定义了一个 Header 结构体，并且实现了 IHeader 接口中的方法
package headerhandler
