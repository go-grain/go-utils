package go_utils

import (
	"github.com/gofrs/uuid"
	"strings"
)

func UID() string {
	uid, _ := uuid.NewV4()
	return strings.ReplaceAll(uid.String(), "-", "")
}
