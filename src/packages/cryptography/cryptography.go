package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Cryptographer interface {
	Hash(data string) (hash string)
	Encrypt(data string, key []byte) (encrypted string)
	Decrypt(ciphertext string, key []byte) (plaintext string, err error)
}

type BasicCryptographer struct{}

func (bc BasicCryptographer) Hash(data string) (hash string) {
	result := sha256.Sum256([]byte(data))
	hashString := hex.EncodeToString(result[:])
	return hashString
}

func (bc BasicCryptographer) Encrypt(data string, key []byte) (encrypted string) {
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	paddedData := padData(data, blockSize)
	message := []byte(paddedData)
	iv := make([]byte, aes.BlockSize)
	cbc := cipher.NewCBCEncrypter(block, iv)
	dataEncrypted := make([]byte, len(message))
	cbc.CryptBlocks(dataEncrypted, message)

	return hex.EncodeToString(dataEncrypted)
}

func padData(data string, blockSize int) []byte {
	paddedData := make([]byte, len(data)+(blockSize-len(data)%blockSize))
	copy(paddedData, []byte(data))
	padLen := blockSize - len(data)%blockSize
	for i := len(data); i < len(paddedData); i++ {
		paddedData[i] = byte(padLen)
	}
	return paddedData
}

func (bc BasicCryptographer) Decrypt(ciphertext string, key []byte) (plaintext string, err error) {
	ciphertextBytes, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := make([]byte, aes.BlockSize)
	cbc := cipher.NewCBCDecrypter(block, iv)

	plaintextBytes := make([]byte, len(ciphertextBytes))
	cbc.CryptBlocks(plaintextBytes, ciphertextBytes)

	plaintextBytes, err = unpadData(plaintextBytes)
	if err != nil {
		return "", err
	}

	return string(plaintextBytes), nil
}

func unpadData(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, nil
	}
	unpadding := int(data[length-1])
	if unpadding > length {
		return nil, fmt.Errorf("Invalid padding")
	}
	return data[:(length - unpadding)], nil
}
