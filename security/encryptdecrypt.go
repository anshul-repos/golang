package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"io"
)

func main() {

	message := "Rmana"
	encrKey := "bCKENV54SyG/pgOy76tjxdgzGREAj7tc8bSBd8/z4RA="
	key, err := base64.StdEncoding.DecodeString(encrKey)
	if err != nil {
		fmt.Println("Error")
	}
	encData, err := encrypt(message, key)
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println("encData", encData)
}

// encrypt uses the same encryption mode (CBC) and padding scheme (PKCS5Padding) for encryption and returns the encrypted message
func encrypt(message string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// The initialization vector is used to ensure that the same plaintext does not encrypt to the same ciphertext when using the same key for encryption.
	// It is important to include the initialization vector as part of the encrypted message, so that the decryption function can use it to correctly decrypt the message
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	encrypter := cipher.NewCBCEncrypter(block, iv)

	// padding
	src := []byte(message)
	blockSize := aes.BlockSize
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	paddedMessage := append(src, padtext...)

	encryptedMessage := make([]byte, len(iv)+len(paddedMessage))

	// copying the iv (initialization vector) to the beginning of the encrypted message.
	copy(encryptedMessage[:aes.BlockSize], iv)
	encrypter.CryptBlocks(encryptedMessage[aes.BlockSize:], paddedMessage)
	return base64.StdEncoding.EncodeToString(encryptedMessage), nil
}

func decrypt(encryptedMessage string, key []byte) (string, error) {
	encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedMessage)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := encryptedBytes[:aes.BlockSize]
	decrypter := cipher.NewCBCDecrypter(block, iv)
	decryptedBytes := make([]byte, len(encryptedBytes)-aes.BlockSize)
	decrypter.CryptBlocks(decryptedBytes, encryptedBytes[aes.BlockSize:])

	paddingLength := int(decryptedBytes[len(decryptedBytes)-1])
	decryptedBytes = decryptedBytes[:len(decryptedBytes)-paddingLength]

	return string(decryptedBytes), nil
}
