package xrsa

import gorsa "github.com/learninto/gorsa"

func applyPriEPubD(data, privateKey string) (string, error) {
	return gorsa.PriKeyEncrypt(data, privateKey)
}
