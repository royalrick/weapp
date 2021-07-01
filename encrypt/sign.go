package encrypt

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"strings"
)

type Signer struct {
	// 是否需要字典排序
	sort bool
	// 需要签名的内容
	parts []string
}

func NewSigner(sort bool, parts ...string) *Signer {
	return &Signer{
		sort:  sort,
		parts: parts,
	}
}

// 对比签名
func (sign *Signer) CompareWith(signature string) bool {
	return signature == sign.Sign()
}

// 生成签名
func (sign *Signer) Sign() string {

	if sign.sort {
		sort.Strings(sign.parts)
	}

	data := strings.Join(sign.parts, "")

	raw := sha1.Sum([]byte(data))

	return hex.EncodeToString(raw[:])
}
