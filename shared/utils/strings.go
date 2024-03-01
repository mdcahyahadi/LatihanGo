package utils

import (
	"encoding/base64"
	"strings"
)

// IsEmptyString func to check empty string
func IsEmptyString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func ByteToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
