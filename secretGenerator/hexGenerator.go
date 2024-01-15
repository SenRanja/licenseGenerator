package secretGenerator

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
)

// 此包用于生产IV值和key值，由于实际使用中IV和key已生成且已固定，此包暂不被main引用

// 加密使用 gcm的分组密码模式
// gcm的 IV值 固定12字节
//secretGenerator.IvGenerator()
// 如果传送密钥长度16B 则aes-128加密 如果密钥长度32B 则aes-256加密，密钥长度不符合将报错
//secretGenerator.CipherKeyGenerator(32)

func generateNonce(ByteLength int) ([]byte, error) {
	nonce := make([]byte, ByteLength)
	_, err := rand.Read(nonce)
	if err != nil {
		return nil, err
	}
	return nonce, nil
}

func IvGenerator() {
	nonce, err := generateNonce(12)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Nonce(iv)'s hex string:", hex.EncodeToString(nonce))
}

func CipherKeyGenerator(ByteLength int) {
	// 仅接受16或32作为int参数
	nonce, err := generateNonce(ByteLength)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Cipher key's hex string:", hex.EncodeToString(nonce))
}
