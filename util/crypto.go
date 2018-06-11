package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"sort"
	"strings"
)

// SignByMD5 多参数通过MD5签名
func SignByMD5(data map[string]string) (string, error) {

	var group []string
	for k, v := range data {
		group = append(group, k+"="+v)
	}

	sort.Strings(group)
	str := strings.Join(group, "&")

	str, err := MD5(str)
	if err != nil {
		return "", err
	}

	return strings.ToUpper(str), nil
}

// MD5 加密
func MD5(str string) (string, error) {
	h := md5.New()
	if _, err := h.Write([]byte(str)); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// PKCS5UnPadding 反补
// Golang AES没有64位的块, 如果采用PKCS5, 那么实质上就是采用PKCS7
func PKCS5UnPadding(plaintext []byte) []byte {
	ln := len(plaintext)

	// 去掉最后一个字节 unPadding 次
	unPadding := int(plaintext[ln-1])
	return plaintext[:(ln - unPadding)]
}

// PKCS5Padding 补全
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

	dSsk, err := base64.StdEncoding.DecodeString(ssk)
	if err != nil {
		return
	}

	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return
	}

	dIv, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return
	}

	block, err := aes.NewCipher(dSsk)
	if err != nil {
		return
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		err = errors.New("cipher too short")
		return
	}

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		err = errors.New("cipher is not a multiple of the block size")
		return
	}

	cipher.NewCBCDecrypter(block, []byte(dIv)).CryptBlocks(ciphertext, ciphertext)

	return PKCS5UnPadding(ciphertext), nil
}
