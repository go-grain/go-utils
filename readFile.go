package go_utils

import "os"

func ReadFile(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return ""
	}
	return string(file)
}
