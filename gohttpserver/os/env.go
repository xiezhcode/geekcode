package os

import "os"

// GetEnvByName 根据环境变量名获取环境变量的值
func GetEnvByName(name string) string {
	return os.Getenv(name)
}
