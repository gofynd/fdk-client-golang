package utils

import "encoding/base64"

func EncodeToBase64(val string) string {
	return base64.StdEncoding.EncodeToString([]byte(val))
}
