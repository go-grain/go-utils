package go_utils

import (
	"fmt"
	"os"
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
)

func GetFileSizeInfo(filePath string) (float64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, fmt.Errorf("无法获取文件信息: %w", err)
	}

	fileSizeInBytes := fileInfo.Size()
	var fileSize float64

	switch {
	case fileSizeInBytes < KB:
		fileSize = float64(fileSizeInBytes)
	case fileSizeInBytes < MB:
		fileSize = float64(fileSizeInBytes) / 1024
	case fileSizeInBytes < GB:
		fileSize = float64(fileSizeInBytes) / (1024 * 1024)
	default:
		fileSize = float64(fileSizeInBytes) / (1024 * 1024 * 1024)
	}

	return fileSize, nil
}

func GetLogFileSize(filePath string) int {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("无法获取文件信息: %s\n", err.Error())
		return 0
	}

	fileSizeInBytes := fileInfo.Size()
	var fileSize float64

	fileSize = float64(fileSizeInBytes / MB)

	return int(fileSize)
}
