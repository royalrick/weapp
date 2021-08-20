package encrypt

import (
	"crypto/sha1"
	"encoding/hex"
	"sort"
	"strings"
)

type Signable struct {
	// 是否需要字典排序
	sort bool
	// 需要签名的内容
	parts []string
}

func NewSignable(sort bool, parts ...string) *Signable {
	return &Signable{
		sort:  sort,
		parts: parts,
	}
}

// 对比签名
func (sign *Signable) IsEqual(signature string) bool {
	return signature == sign.Sign()
}

// 生成签名
func (sign *Signable) Sign() string {

	if sign.sort {
		sort.Strings(sign.parts)
	}

	data := strings.Join(sign.parts, "")

	raw := sha1.Sum([]byte(data))

	return hex.EncodeToString(raw[:])
}
