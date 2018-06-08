package util

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b: b,
		// AES 算法有 AES-128、AES-192、AES-256三种
		// 分别对应的key是 16、24、32字节长度
		// 同样对应的加密解密区块长度BlockSize也是16、24、32字节长度。
		blockSize: b.BlockSize(),
	}
}

// NewECBEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return (*ecb)(newECB(b))
}

func (e *ecb) BlockSize() int { return e.blockSize }

func (e *ecb) CryptBlocks(dst, src []byte) {
	if len(src)%e.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		e.b.Encrypt(dst, src[:e.blockSize])
		src = src[e.blockSize:]
		dst = dst[e.blockSize:]
	}
}

// NewECBDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return (*ecb)(newECB(b))
}

// AesECBDecrypt 解密
// @key 商户号密钥
func AesECBDecrypt(cipher []byte, key string) (plaintext []byte, err error) {
	if len(cipher) < aes.BlockSize {
		return nil, errors.New("cipher too short")
	}
	// ECB mode always works in whole blocks.
	if len(cipher)%aes.BlockSize != 0 {
		return nil, errors.New("cipher is not a multiple of the block size")
	}

	key, err = MD5(key)
	if err != nil {
		return
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return
	}

	NewECBDecrypter(block).CryptBlocks(cipher, cipher)

	return PKCS5UnPadding(cipher), nil
}

// PKCS5UnPadding unpadding by PKCS5
// Golang AES没有64位的块, 如果采用PKCS5, 那么实质上就是采用PKCS7
func PKCS5UnPadding(plaintext []byte) []byte {
	ln := len(plaintext)

	// 去掉最后一个字节 unPadding 次
	unPadding := int(plaintext[ln-1])
	return plaintext[:(ln - unPadding)]
}
