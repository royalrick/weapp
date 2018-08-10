// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Electronic Code Book (ECB) mode.

// ECB provides confidentiality by assigning a fixed ciphertext block to each
// plaintext block.

// See NIST SP 800-38A, pp 08-09

package ecb

import (
	"crypto/cipher"
)

type ecb struct {
	b         cipher.Block
	blockSize int
}

func newECB(b cipher.Block) *ecb {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

type Encrypter ecb

// NewEncrypter returns a BlockMode which encrypts in electronic code book
// mode, using the given Block.
func NewEncrypter(b cipher.Block) cipher.BlockMode {
	return (*Encrypter)(newECB(b))
}

func (x *Encrypter) BlockSize() int { return x.blockSize }

func (x *Encrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Encrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}

type Decrypter ecb

// NewDecrypter returns a BlockMode which decrypts in electronic code book
// mode, using the given Block.
func NewDecrypter(b cipher.Block) cipher.BlockMode {
	return (*Decrypter)(newECB(b))
}

func (x *Decrypter) BlockSize() int { return x.blockSize }

func (x *Decrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
