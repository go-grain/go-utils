package go_utils

import (
	"os"
	"strings"
)

func PathIsNotExist(path string) bool {
	f, err := os.Stat(path)
	if os.IsNotExist(err) {
		return true
	} else {
		if f.IsDir() {
			return false
		}
		return true
	}
}

func FileIsNotExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}

// ext 获取文件后缀
func Ext(filename string) string {
	index := strings.LastIndex(filename, ".")
	if index == -1 || index == len(filename)-1 {
		return ""
	}
	return strings.TrimSpace(filename[index+1:])
}

// FileNameWithoutExt 获取文件名
func FileNameWithoutExt(fileName, suffix string) string {
	fileName = strings.TrimSuffix(fileName, suffix)
	return fileName
}
