package utils

import (
	"encoding/base64"
	"net/url"
)

//EncodeToBase64 gives base64 encoded string
func EncodeToBase64(val string) string {
	return base64.StdEncoding.EncodeToString([]byte(val))
}

//MapToURLEncodedString converts map to url encoded string
func MapToURLEncodedString(valMap map[string]string) (encodedData string) {
	data := url.Values{}
	for k, v := range valMap {
		data.Set(k, v)
	}
	encodedData = data.Encode()
	return
}
