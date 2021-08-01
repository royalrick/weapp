package encrypt

import "bytes"

const pkcs7blocksize = 32

// 对需要加密的明文进行填充补位
// plaintext 需要进行填充补位操作的明文
// 返回补齐明文字符串
func pkcs7padding(plaintext []byte) []byte {
	//计算需要填充的位数
	pad := pkcs7blocksize - len(plaintext)%pkcs7blocksize
	if pad == 0 {
		pad = pkcs7blocksize
	}

	//获得补位所用的字符
	text := bytes.Repeat([]byte{byte(pad)}, pad)

	return append(plaintext, text...)

}

// 对解密后的明文进行补位删除
// plaintext 解密后的明文
// 返回删除填充补位后的明文和
func pkcs7unPadding(plaintext []byte) []byte {
	ln := len(plaintext)

	// 获取最后一个字符的 ASCII
	pad := int(plaintext[ln-1])
	if pad < 1 || pad > pkcs7blocksize {
		pad = 0
	}

	return plaintext[:(ln - pad)]
}
