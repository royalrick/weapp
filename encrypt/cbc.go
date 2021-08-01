package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

type CBC struct {
	iv      []byte
	key     []byte
	content []byte
}

func NewCBC(iv, key, content []byte) *CBC {
	return &CBC{
		iv:      iv,
		key:     key,
		content: content,
	}
}

// 加密数据
func (cbc *CBC) Encrypt() ([]byte, error) {

	cbc.content = pkcs7padding(cbc.content)

	if len(cbc.content)%aes.BlockSize != 0 {
		return nil, errors.New("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(cbc.key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(cbc.content))
	cbc.iv = cbc.iv[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, cbc.iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, cbc.iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], cbc.content)

	return ciphertext, nil
}

// CBC解密数据
func (cbc *CBC) Decrypt() ([]byte, error) {
	block, err := aes.NewCipher(cbc.key)
	if err != nil {
		return nil, err
	}

	size := aes.BlockSize
	cbc.iv = cbc.iv[:size]

	if len(cbc.content) < size {
		return nil, errors.New("ciphertext too short")
	}

	if len(cbc.content)%size != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, cbc.iv)
	mode.CryptBlocks(cbc.content, cbc.content)

	return pkcs7unPadding(cbc.content), nil
}
