package utils

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	content := "123456"
	md5 := Md5(content)

	if "E10ADC3949BA59ABBE56E057F20F883E" != md5 {
		t.Errorf("md5 failed")
	}
}

func TestSHA1(t *testing.T) {
	content := "123456"
	sha1 := fmt.Sprintf("%x", SHA1([]byte(content)))

	if "7c4a8d09ca3762af61e59520943dc26494f8941b" != sha1 {
		t.Errorf("sha1 failed")
	}
}
