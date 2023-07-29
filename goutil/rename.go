package goutil

import (
	"fmt"
	"github.com/go-pay/gopay/pkg/xlog"
	"os"
	"strings"
)

// 修改文件后缀
func RenameAllDirFileSuffix(path, oldSuffix, newSuffix string) {
	ReadDir(path)
	for _, file := range DirList {
		if FileExtension(file.Filename) == oldSuffix {
			err := os.Rename(fmt.Sprintf("%s/%s", file.Filepath, file.Filename), fmt.Sprintf("%s/%s%s", file.Filepath, FileNameWithoutExtension(file.Filename, oldSuffix), newSuffix))
			//err := os.Rename(fmt.Sprintf("%s/%s", file.Filepath, file.Filename), fmt.Sprintf("%s/%s", file.Filepath, strings.ReplaceAll(file.Filename, "..", ".")))
			if err != nil {
				xlog.Info(err)
				return
			}
		}
	}
	xlog.Info("success")
}

func RenameFileNameSuffix(filename, oldSuffix, newSuffix string) {
	newSuffix = strings.TrimSpace(newSuffix)
	oldSuffix = strings.TrimSpace(oldSuffix)
	fmt.Println(FileNameWithoutExtension(filename, oldSuffix) + newSuffix)
	err := os.Rename(filename, FileNameWithoutExtension(filename, oldSuffix)+newSuffix)
	if err != nil {
		xlog.Info(err)
		return
	}
}
