package go_utils

import (
	"strconv"
	"strings"
)

func ToLower(val string) string {
	if len(val) > 0 {
		v := strings.ToLower(val[:1])
		val = v + val[1:]
	}
	return val
}

func ToTitle(val string) string {
	if len(val) > 0 {
		v := strings.ToTitle(val[:1])
		val = v + val[1:]
	}
	return val
}

func ToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	}
	return ""
}
