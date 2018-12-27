package functionbz

import (
	"os"
)

// GetRunPath 获取当天运行的路径
func GetRunPath() (path string, err error) {
	path, err = os.Getwd()
	// log.Println(path)
	return
}
