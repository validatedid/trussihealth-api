package cryptography_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/validatedid/trussihealth-api/src/packages/cryptography"
)

func TestHashData(t *testing.T) {
	basicCryptography := cryptography.BasicCryptographer{}

	dataHashed := basicCryptography.Hash("Hello")
	expectedHash := "185f8db32271fe25f561a6fc938b2e264306ec304eda518007d1764826381969"

	assert.Equal(t, expectedHash, dataHashed, "The two hases should be the same.")

}

func TestEncryptData(t *testing.T) {

	basicCryptography := cryptography.BasicCryptographer{}

	data := "Hello, world!"
	// message := []byte(data)

	key := []byte("0123456789abcdef0123456789abcdef")

	// block, err := aes.NewCipher(key)
	// if err != nil {
	// 	panic(err)
	// }

	// iv := make([]byte, aes.BlockSize)

	// cbc := cipher.NewCBCEncrypter(block, iv)

	// expectedDataEncrypted := make([]byte, len(message))
	// cbc.CryptBlocks(expectedDataEncrypted, message)

	dataEncrypted := basicCryptography.Encrypt(data, key)

	assert.Equal(t, string("expectedDataEncrypted"), dataEncrypted, "The two encrypted data should be the same.")

}
