package logger

import (
	"log"
	"os"
)

var Logger *log.Logger

// 初始化 logger
func init() {
	// 创建 Logger，设置日志输出格式
	Logger = log.New(os.Stdout, "", log.Lshortfile|log.Ldate|log.Ltime)
}
