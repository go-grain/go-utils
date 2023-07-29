package goutil

import (
	"fmt"
	"github.com/go-pay/gopay/pkg/xlog"
	"os"
)

type File struct {
	Filename string
	Filepath string
}

var DirList []File

func ReadDir(path string) []File {
	//core.StartGrain()
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil
	}
	for _, entry := range dir {
		if entry.IsDir() {
			ReadDir(path + "/" + entry.Name())
		} else {
			DirList = append(DirList, File{Filename: entry.Name(), Filepath: path})
		}
	}
	return DirList
}

func FmtProjectCode(projectPath string) error {
	ReadDir(projectPath)
	for _, file := range DirList {
		if FileExtension(file.Filename) == "go" {
			err := FmtCode(fmt.Sprintf("%s/%s", file.Filepath, file.Filename))
			if err != nil {
				xlog.Info(err)
				return err
			}
		}
	}
	xlog.Info("success")
	return nil
}
