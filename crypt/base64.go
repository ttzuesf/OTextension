package crypt

import (
	"encoding/base64"
)

func Base64Encoding(dst []byte) string {
	return base64.StdEncoding.EncodeToString(dst)
}

func Base64Decoding(dst string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(dst)
}
