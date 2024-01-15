package gcmAesCipher

import (
	"reflect"
	"testing"
)

func Test_encryptGCM(t *testing.T) {
	type args struct {
		plaintext string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Cipher: wuwuwuhahaha",
			args: args{
				plaintext: "wuwuwuhahaha",
			},
		}, {
			name: "Cipher: heihei",
			args: args{
				plaintext: "heihei",
			},
		}, {
			name: "Cipher: qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq",
			args: args{
				plaintext: "qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cipher := EncryptGCM([]byte(tt.args.plaintext))
			//cipherHex := hex.EncodeToString(cipher)
			//fmt.Println("密文的byte类型", cipher)
			//fmt.Println("密文的hex编码", cipherHex)
			plaintext := DecryptGCM(cipher)
			//fmt.Println("明文的byte类型", plaintext)
			//fmt.Println("明文字符串", string(plaintext))
			if !reflect.DeepEqual([]byte(tt.args.plaintext), plaintext) {
				t.Fail()
			}
		})
	}
}
