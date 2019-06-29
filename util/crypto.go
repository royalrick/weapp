package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
)

// PKCS5UnPadding 反补
// Golang AES没有64位的块, 如果采用PKCS5, 那么实质上就是采用PKCS7
func PKCS5UnPadding(plaintext []byte) ([]byte, error) {
	ln := len(plaintext)

	// 去掉最后一个字节 unPadding 次
	unPadding := int(plaintext[ln-1])

	if unPadding > ln {
		return []byte{}, errors.New("数据不正确")
	}

	return plaintext[:(ln - unPadding)], nil
}

// PKCS5Padding 补位
// Golang AES没有64位的块, 如果采用PKCS5, 那么实质上就是采用PKCS7
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// Validate 对数据包进行签名校验，确保数据的完整性。
//
// @rawData 不包括敏感信息的原始数据字符串，用于计算签名。
// @signature 使用 sha1( rawData + sessionkey ) 得到字符串，用于校验用户信息
// @ssk 微信 session_key
func Validate(rawData, ssk, signature string) bool {
	r := sha1.Sum([]byte(rawData + ssk))

	return signature == hex.EncodeToString(r[:])
}

// CBCDecrypt CBC解密数据
//
// @ssk 通过 Login 向微信服务端请求得到的 session_key
// @data 小程序通过 api 得到的加密数据(encryptedData)
// @iv 小程序通过 api 得到的初始向量(iv)
func CBCDecrypt(ssk, data, iv string) (bts []byte, err error) {
	key, err := base64.StdEncoding.DecodeString(ssk)
	if err != nil {
		return
	}

	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return
	}

	rawIV, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	size := aes.BlockSize

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < size {
		err = errors.New("cipher too short")
		return
	}

	// CBC mode always works in whole blocks.
	if len(ciphertext)%size != 0 {
		err = errors.New("cipher is not a multiple of the block size")
		return
	}

	mode := cipher.NewCBCDecrypter(block, rawIV[:size])
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	return PKCS5UnPadding(plaintext)
}
