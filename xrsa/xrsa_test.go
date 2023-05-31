package xrsa

import (
	"fmt"
	"testing"
)

var Pubkey = `-----BEGIN Public key-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCHokoHFmvJIhKjM76ldAU3fDBgIuX8NFjMVV7+fFYISHPbd9HMxeA9vPXHvtGiYfw0dmSFqwcxA463WtO7XQQ1Nu5HE54id09g1TzQDhXKVVANtDiKJNAUQWCuwHPIqiXHCQf6v01ynR8jR2z40OpMKe1sdnEAloeY9q1l181vbwIDAQAB
-----END Public key-----
`

var Pirvatekey = `-----BEGIN Private key-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAIeiSgcWa8kiEqMzvqV0BTd8MGAi5fw0WMxVXv58VghIc9t30czF4D289ce+0aJh/DR2ZIWrBzEDjrda07tdBDU27kcTniJ3T2DVPNAOFcpVUA20OIok0BRBYK7Ac8iqJccJB/q/TXKdHyNHbPjQ6kwp7Wx2cQCWh5j2rWXXzW9vAgMBAAECgYB6K9yydaexDFftWXaoYdExIVQRxF2UxzIVG/DtGeIEo/53+X2pDbPm6IYa3e7GbaxXNS1mmZ9oruOmlNGTOz3FwWJN+C33WQMLBOlO2gUQSqLZ4X8EvTEBiMvF42U5pDqyAh0EgGtDL4TtA34CgEyX7Iw7rKgfhVvE1OWuDiQ3QQJBAL7rQZt0RARAEqG+WAEPWaN3iUsLkiClZnyfF3hVQug3n/OVSDm7iuKZyYRHIEanTGrfabq7fPw/qD+4fJcMj00CQQC13obu3bhBwCc01ApWR6cy4+liKexpGTvUK+5LpPiM72jFhPyRXdmYsdeXmE6SuxIYDx6xJVip4O1YuC99YxOrAkBjDkas9Gbx2ZiRKOQaMK+ue6/VKvy3SXniQNz5hys+ttWbmSGvKpoFtgrzQcACSH0CmkYOJ4bSjeiqnvqtmEulAkEAkjtdtTxzhfKR06lWsm8kogedRP++hfbzIzM7hHkeHHv3ezHlvqB+cIc2eT7Olq5x6wRlQjxsIROo47gc/y2lxwJAazp24jLnV9kGT6+hC8k79z9kFqs9x8H7Vil1IF/ciwSMSHq4IlT2X8CkgyR6iAZd9SNU24tcj05uVbpP14zC+g==
-----END Private key-----
`

func Test_PriKeyEncrypt(t *testing.T) {
	res, err := GenerateKey(1024)
	if err != nil {
		fmt.Println(err)
		return
	}
	publicKey := res.PublicKeyBase64
	privateKey := res.PrivateKeyBase64

	fmt.Println("\n私钥: \n\r" + privateKey)
	fmt.Println("\n公钥: \n\r" + publicKey)
	fmt.Println("\n私钥加密 —— 公钥解密")

	str := `{"domainId":"id", "externalUserId":"test001"}`
	fmt.Println("\n\r明文：\r\n" + str)
	encodedData, err := PriKeyEncrypt(str, privateKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n密文：\r\n" + encodedData)

	decodedData, err := PublicDecrypt(encodedData, publicKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\n解密后文字: \r\n" + decodedData)
}
