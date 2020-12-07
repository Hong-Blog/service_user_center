package utils

import "testing"

func TestPasswordEncrypt(t *testing.T) {
	user := "root"
	encryptedPassword := "CGUx1FN++xS+4wNDFeN6DA=="

	encrypt := PasswordEncrypt("123456", user)

	if encrypt != encryptedPassword {
		t.Errorf("%s is not %s", encrypt, encryptedPassword)
	}
}

func TestPasswordDecrypt(t *testing.T) {
	user := "root"
	encryptedPassword := "CGUx1FN++xS+4wNDFeN6DA=="
	password := "123456"

	decrypt := PasswordDecrypt(encryptedPassword, user)

	if decrypt != password {
		t.Errorf("%s is not %s", decrypt, password)
	}
}
