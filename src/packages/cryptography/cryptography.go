package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
)

type Cryptographer interface {
	Hash()
	Encrypt()
}

type BasicCryptographer struct{}

func (bc BasicCryptographer) Hash(data string) (hash string) {
	result := sha256.Sum256([]byte(data))
	hashString := hex.EncodeToString(result[:])
	return hashString
}

func (bc BasicCryptographer) Encrypt(data string, key []byte) (encrypted string) {
	message := []byte(data)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	iv := make([]byte, aes.BlockSize)

	cbc := cipher.NewCBCEncrypter(block, iv)

	dataEncrypted := make([]byte, len(message))
	cbc.CryptBlocks(dataEncrypted, message)

	return string(dataEncrypted)
}
