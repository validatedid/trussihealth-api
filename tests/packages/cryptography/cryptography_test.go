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
	data := "This message is large enough."
	key := []byte("thisis32bitlongpassphraseimusing")
	expectedDataEncrypted := "5bc234eb44fee4ea5d6004dfda23cf824d49a20fd90a88be6c21dccb1d4ad09e"
	dataEncrypted := basicCryptography.Encrypt(data, key)
	assert.Equal(t, expectedDataEncrypted, dataEncrypted, "The two encrypted data should be the same.")
}

func TestEncryptDataShortMessage(t *testing.T) {
	basicCryptography := cryptography.BasicCryptographer{}
	data := "Hello World!"
	key := []byte("thisis32bitlongpassphraseimusing")
	expectedDataEncrypted := "cbbb4023658a770a5a053b5cf92f4482"
	dataEncrypted := basicCryptography.Encrypt(data, key)
	assert.Equal(t, expectedDataEncrypted, dataEncrypted, "The two encrypted data should be the same.")
}

func TestDecryptData(t *testing.T) {
	basicCryptography := cryptography.BasicCryptographer{}
	key := []byte("thisis32bitlongpassphraseimusing")
	encryptedMessage := "5bc234eb44fee4ea5d6004dfda23cf824d49a20fd90a88be6c21dccb1d4ad09e"
	dataDecrypted, _ := basicCryptography.Decrypt(encryptedMessage, key)
	expectedMessage := "This message is large enough."
	assert.Equal(t, expectedMessage, dataDecrypted, "The data decrypted is not the expected")
}


func TestDecryptShortMessage(t *testing.T) {
	basicCryptography := cryptography.BasicCryptographer{}
	key := []byte("thisis32bitlongpassphraseimusing")
	encryptedMessage := "cbbb4023658a770a5a053b5cf92f4482"
	dataDecrypted, _ := basicCryptography.Decrypt(encryptedMessage, key)
	expectedMessage := "Hello World!"
	assert.Equal(t, expectedMessage, dataDecrypted, "The data decrypted is not the expected")
}
