package gcmAesCipher

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"
)

// GCM的分组密码模式不需要指定 padding，因为本身支持不定长分组

var hex_cipher_key = "fd2d9fb0bb4bc7a7e5d36e0cce8d26573950c4461c4d29afd46214e37e686197"
var hex_iv = "29cd7489b0a30e28fb589377"

func EncryptGCM(plainBytes []byte) []byte {

	cipher_key, _ := hex.DecodeString(hex_cipher_key)
	cipher_iv, _ := hex.DecodeString(hex_iv)

	block, _ := aes.NewCipher(cipher_key)

	aesGCM, _ := cipher.NewGCM(block)

	ciphertext := aesGCM.Seal(nil, cipher_iv, plainBytes, nil)

	// 提取密文部分，返回纯密文
	// 生成的ciphertext默认直接拼接了aesGCM加密的digest部分（有认证功能），此处不想携带digest部分，故切割出 encryptedData
	// 由于GCM解密时，必须有认证部分，因此建议密文无需切割digest部分
	//encryptedData := ciphertext[:len(ciphertext)-aesGCM.Overhead()]

	return ciphertext
}

func DecryptGCM(cipherBytes []byte) []byte {
	// 解密函数必须传入 加密时的digest用作认证，否则直接报错

	cipher_key, _ := hex.DecodeString(hex_cipher_key)
	cipher_iv, _ := hex.DecodeString(hex_iv)

	block, _ := aes.NewCipher(cipher_key)

	aesGCM, _ := cipher.NewGCM(block)

	plaintext, err := aesGCM.Open(nil, cipher_iv, cipherBytes, nil)
	if err != nil {
		log.Fatal(err)
	}

	return plaintext
}
