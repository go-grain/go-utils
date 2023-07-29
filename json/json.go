package xjson

import (
	"errors"
	"github.com/bytedance/sonic"
	"log"
)

type G map[string]any

func Marshal(data any) []byte {
	marshal, err := sonic.Marshal(data)
	if err != nil {
		log.Print(err)
		return nil
	}
	return marshal
}

func Unmarshal(data []byte, v any) error {
	if len(data) == 0 {
		return errors.New("value is empty")
	}
	err := sonic.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
}
