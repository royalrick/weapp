package weapp

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
)

const pkcs7blocksize = 32

// pkcs7encode 对需要加密的明文进行填充补位
// @plaintext 需要进行填充补位操作的明文
// 返回补齐明文字符串
func pkcs7encode(plaintext []byte) []byte {
	//计算需要填充的位数
	pad := pkcs7blocksize - len(plaintext)%pkcs7blocksize
	if pad == 0 {
		pad = pkcs7blocksize
	}

	//获得补位所用的字符
	text := bytes.Repeat([]byte{byte(pad)}, pad)

	return append(plaintext, text...)
}

// pkcs7decode 对解密后的明文进行补位删除
// @plaintext 解密后的明文
// 返回删除填充补位后的明文和
func pkcs7decode(plaintext []byte) []byte {
	ln := len(plaintext)

	// 获取最后一个字符的 ASCII
	pad := int(plaintext[ln-1])
	if pad < 1 || pad > pkcs7blocksize {
		pad = 0
	}

	return plaintext[:(ln - pad)]
}

// ValidateSignature 对数据包进行签名校验，确保数据的完整性。
//
// @rawData 不包括敏感信息的原始数据字符串，用于计算签名。
// @signature 使用 sha1( rawData + sessionkey ) 得到字符串，用于校验用户信息
// @ssk 微信 session_key
func ValidateSignature(rawData, ssk, signature string) bool {
	r := sha1.Sum([]byte(rawData + ssk))

	return signature == hex.EncodeToString(r[:])
}

// cbcEncrypt CBC 加密数据
func cbcEncrypt(key, plaintext, iv []byte) ([]byte, error) {
	if len(plaintext)%aes.BlockSize != 0 {
		return nil, errors.New("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv = iv[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

// CBCDecrypt CBC解密数据
//
// @ssk 通过 Login 向微信服务端请求得到的 session_key
// @ciphertext 小程序通过 api 得到的加密数据(encryptedData)
// @iv 小程序通过 api 得到的初始向量(iv)
func cbcDecrypt(key, ciphertext, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	size := aes.BlockSize
	iv = iv[:size]
	ciphertext = ciphertext[aes.BlockSize:]

	if len(ciphertext) < size {
		return nil, errors.New("ciphertext too short")
	}

	if len(ciphertext)%size != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(ciphertext, ciphertext)

	return pkcs7decode(plaintext), nil
}

// decryptShareData CBC解密数据
//
// @ssk 通过 Login 向微信服务端请求得到的 session_key
// @ciphertext 小程序通过 api 得到的加密数据(encryptedData)
// @iv 小程序通过 api 得到的初始向量(iv)
func decryptShareData(ssk, ciphertext, iv string) ([]byte, error) {
	key, err := base64.StdEncoding.DecodeString(ssk)
	if err != nil {
		return nil, err
	}

	cipher, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	rawIV, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}

	return cbcDecrypt(key, cipher, rawIV)
}
