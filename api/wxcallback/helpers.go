package wxcallback

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"strings"
)

// GetSignature 获取签名
func GetSignature(timestamp, nonce string, encrypted string, token string) string {
	data := []string{
		encrypted,
		token,
		timestamp,
		nonce,
	}
	sort.Strings(data)
	s := sha1.New()
	_, err := io.WriteString(s, strings.Join(data, ""))
	if err != nil {
		panic("签名错误：sign error")
	}
	return fmt.Sprintf("%x", s.Sum(nil))
}

func makeRandomString(length int) string {
	randStr := ""
	strSource := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyl"
	maxLength := len(strSource) - 1
	for i := 0; i < length; i++ {
		randomNum := rand.Intn(maxLength)
		randStr += strSource[randomNum : randomNum+1]
	}
	return randStr
}

func pKCS7Pad(plainText []byte, blockSize int) []byte {
	// block size must be bigger or equal 2
	if blockSize < 1<<1 {
		panic("block size is too small (minimum is 2 bytes)")
	}
	// block size up to 255 requires 1 byte padding
	if blockSize < 1<<8 {
		// calculate padding length
		padLen := padLength(len(plainText), blockSize)

		// define PKCS7 padding block
		padding := bytes.Repeat([]byte{byte(padLen)}, padLen)

		// apply padding
		padded := append(plainText, padding...)
		return padded
	}
	// block size bigger or equal 256 is not currently supported
	panic("unsupported block size")
}

func padLength(sliceLength, blockSize int) int {
	padLen := blockSize - sliceLength%blockSize
	if padLen == 0 {
		padLen = blockSize
	}
	return padLen
}
