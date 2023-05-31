package xrsa

import "github.com/learninto/gorsa"

func PriKeyEncrypt(data, privateKey string) (string, error) {
	return gorsa.PriKeyEncrypt(data, privateKey)
}

func PriKeyDecrypt(data, privateKey string) (string, error) {
	return gorsa.PriKeyDecrypt(data, privateKey)
}

func PublicEncrypt(data, publicKey string) (string, error) {
	return gorsa.PublicEncrypt(data, publicKey)
}

func PublicDecrypt(data, publicKey string) (string, error) {
	return gorsa.PublicDecrypt(data, publicKey)
}

func GenerateKey(bits int) (resp gorsa.RSAKey, err error) {
	return gorsa.GenerateKey(bits)
}
