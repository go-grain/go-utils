package goutil

import (
	"fmt"
	"github.com/go-pay/gopay/pkg/xlog"
	"os"
	"strings"
)

// ReplaceAllDirFileContent 替换xx目录所有文件内容
func ReplaceAllDirFileContent(path, old, new, suffix string) {
	ReadDir(path)
	for _, file := range DirList {
		if FileExtension(file.Filename) == suffix {
			bytes, err := os.ReadFile(fmt.Sprintf("%s/%s", file.Filepath, file.Filename))
			if err != nil {
				xlog.Info(err)
				return
			}
			newStr := strings.ReplaceAll(string(bytes), old, new)
			err = os.WriteFile(fmt.Sprintf("%s/%s", file.Filepath, file.Filename), []byte(newStr), 0644)
			if err != nil {
				xlog.Info(err)
				return
			}
			xlog.Info(file.Filename, file.Filepath+file.Filename+" success")
		}
	}
}

func ReplaceFileContent(filename, old, new string) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		xlog.Info(err)
		return
	}
	newStr := strings.ReplaceAll(string(bytes), old, new)
	err = os.WriteFile(filename, []byte(newStr), 0644)
	if err != nil {
		xlog.Info(err)
		return
	}
	fmt.Println("success")
}
