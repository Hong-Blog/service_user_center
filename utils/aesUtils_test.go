package utils

import (
	"testing"
)

const key = "asdfzxcv"

func TestAesDecrypt(t *testing.T) {
	encryptedString := "D+RL6Hhbp2TrlFnlsu3Jtg=="
	decryptedString := "123qwe"

	decrypt := AesDecrypt(encryptedString, key)
	if decryptedString != decrypt {
		t.Errorf("encrypt is should be %s but is %s", decryptedString, decrypt)
	}
}

func TestAesEncrypt(t *testing.T) {
	str := "123qwe"
	encrypt := AesEncrypt(str, key)
	encryptString := "D+RL6Hhbp2TrlFnlsu3Jtg=="

	if encryptString != encrypt {
		t.Errorf("encrypt is should be %s but is %s", encryptString, encrypt)
	}
}
