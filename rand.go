package go_utils

import (
	"math/rand"
	"time"
)

func RandomString(length int) int64 {
	rand.NewSource(time.Now().UnixNano())
	const charset = "0123456789"
	var result []byte
	for i := 0; i < length; i++ {
		result = append(result, charset[rand.Intn(len(charset))])
	}
	return String2Int64(string(result))
}

func RandomCharset(length int) string {
	rand.NewSource(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var result []byte
	for i := 0; i < length; i++ {
		result = append(result, charset[rand.Intn(len(charset))])
	}
	return string(result)
}
