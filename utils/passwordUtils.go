package utils

const securityKey = "929123f8f17944e8b0a531045453e1f1"

func PasswordEncrypt(password string, salt string) string {
	hash := Md5(salt + securityKey)
	encrypt := AesEncrypt(password, hash)
	return encrypt
}

func PasswordDecrypt(encryptedPassword string, salt string) string {
	hash := Md5(salt + securityKey)
	decrypt := AesDecrypt(encryptedPassword, hash)
	return decrypt
}
