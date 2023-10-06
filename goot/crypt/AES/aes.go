package AES

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"github.com/andreburgaud/crypt2go/padding"
)

func CBCEncrypt(plaintext []byte, key []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	padder := padding.NewPkcs7Padding(aes.BlockSize)
	m, err := padder.Pad(plaintext)
	if err != nil {
		return nil, errors.New("wrong padding process")
	}
	// padding using PKCS7
	ciphertext := make([]byte, aes.BlockSize+len(m))
	iv := ciphertext[:aes.BlockSize]
	if _, err = rand.Read(iv); err != nil { // 将同时写到 ciphertext 的开头
		return nil, errors.New("add iv wrong")
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], m)
	return ciphertext, nil
}

// CBCDecrypt AES-CBC
func CBCDecrypt(ciphertext []byte, key []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("creat cipher method wrong")
	}

	iv := ciphertext[:aes.BlockSize] // extract iv from ciphertext
	c := ciphertext[aes.BlockSize:]  // the ciphertext context
	plaintext := make([]byte, len(c))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, c)

	padder := padding.NewPkcs7Padding(aes.BlockSize) // remove padding from msg
	message, err := padder.Unpad(plaintext)
	if err != nil {
		return nil, errors.New("unpad wrong")
	}
	return message, nil
}
