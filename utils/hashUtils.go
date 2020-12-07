package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%X", has)
}

func SHA1(data []byte) []byte {
	h := sha1.New()
	h.Write(data)
	return h.Sum(nil)
}
