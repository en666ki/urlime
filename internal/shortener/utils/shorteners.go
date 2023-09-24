package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5Shortener(data []byte) string {
	sum := md5.Sum([]byte(data))
	var s [md5.Size]byte
	for i, b := range sum {
		s[i] = fmt.Sprintf("%x", b)[0]
	}
	return string(s[:])
}
