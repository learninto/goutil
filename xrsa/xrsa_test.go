package xrsa

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/learninto/gorsa"
)

var Pubkey = `-----BEGIN Public key-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCHokoHFmvJIhKjM76ldAU3fDBgIuX8NFjMVV7+fFYISHPbd9HMxeA9vPXHvtGiYfw0dmSFqwcxA463WtO7XQQ1Nu5HE54id09g1TzQDhXKVVANtDiKJNAUQWCuwHPIqiXHCQf6v01ynR8jR2z40OpMKe1sdnEAloeY9q1l181vbwIDAQAB
-----END Public key-----
`

var Pirvatekey = `-----BEGIN Private key-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAIeiSgcWa8kiEqMzvqV0BTd8MGAi5fw0WMxVXv58VghIc9t30czF4D289ce+0aJh/DR2ZIWrBzEDjrda07tdBDU27kcTniJ3T2DVPNAOFcpVUA20OIok0BRBYK7Ac8iqJccJB/q/TXKdHyNHbPjQ6kwp7Wx2cQCWh5j2rWXXzW9vAgMBAAECgYB6K9yydaexDFftWXaoYdExIVQRxF2UxzIVG/DtGeIEo/53+X2pDbPm6IYa3e7GbaxXNS1mmZ9oruOmlNGTOz3FwWJN+C33WQMLBOlO2gUQSqLZ4X8EvTEBiMvF42U5pDqyAh0EgGtDL4TtA34CgEyX7Iw7rKgfhVvE1OWuDiQ3QQJBAL7rQZt0RARAEqG+WAEPWaN3iUsLkiClZnyfF3hVQug3n/OVSDm7iuKZyYRHIEanTGrfabq7fPw/qD+4fJcMj00CQQC13obu3bhBwCc01ApWR6cy4+liKexpGTvUK+5LpPiM72jFhPyRXdmYsdeXmE6SuxIYDx6xJVip4O1YuC99YxOrAkBjDkas9Gbx2ZiRKOQaMK+ue6/VKvy3SXniQNz5hys+ttWbmSGvKpoFtgrzQcACSH0CmkYOJ4bSjeiqnvqtmEulAkEAkjtdtTxzhfKR06lWsm8kogedRP++hfbzIzM7hHkeHHv3ezHlvqB+cIc2eT7Olq5x6wRlQjxsIROo47gc/y2lxwJAazp24jLnV9kGT6+hC8k79z9kFqs9x8H7Vil1IF/ciwSMSHq4IlT2X8CkgyR6iAZd9SNU24tcj05uVbpP14zC+g==
-----END Private key-----
`

func Test_applyPriEPubD(t *testing.T) {
	text := `{"domainId":"guanbi", "externalUserId":"test001"}`
	prienctypt, err := gorsa.PriKeyEncrypt(text, Pirvatekey)
	if err != nil {
		return
	}
	//println(`{"domainId":"domId","externalUserId":"UID","timestamp":1521616977}`)
	fmt.Printf("http://databi.zzgqsh.com/m/app/v082b65dd1b564e04b0ff563?provider=guanbi&ssoToken=%s", hex.EncodeToString([]byte(prienctypt)))
	//println(prienctypt)

	//pubdecrypt, err := gorsa.PublicDecrypt(prienctypt, Pubkey)
	//if err != nil {
	//	return err
	//}
	//println(pubdecrypt)
	//if string(pubdecrypt) != `liyong` {
	//	return errors.New(`Decryption failed`)
	//}

	type args struct {
		data       string
		privateKey string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := applyPriEPubD(tt.args.data, tt.args.privateKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("applyPriEPubD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("applyPriEPubD() got = %v, want %v", got, tt.want)
			}
		})
	}
}
