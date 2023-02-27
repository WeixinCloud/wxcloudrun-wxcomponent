package wxcallback

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

type prpCrypt struct {
	Key []byte
	Iv  []byte
}

func NewPrpCrypt(aesKey string) *prpCrypt {
	instance := new(prpCrypt)
	//网络字节序
	instance.Key, _ = base64.StdEncoding.DecodeString(aesKey + "=")
	instance.Iv = randomIv()
	return instance
}

func randomIv() []byte {
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic("random iv error")
	}
	return iv
}

func (prp *prpCrypt) decrypt(encrypted string) ([]byte, error) {
	encryptedBytes, _ := base64.StdEncoding.DecodeString(encrypted)
	k := len(prp.Key) //PKCS#7
	if len(encryptedBytes)%k != 0 {
		panic("ciphertext size is not multiple of aes key length")
	}
	block, err := aes.NewCipher(prp.Key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, prp.Iv)
	plainText := make([]byte, len(encryptedBytes))
	blockMode.CryptBlocks(plainText, encryptedBytes)
	return plainText, nil
}

func (prp *prpCrypt) encrypt(plainText []byte) ([]byte, error) {
	k := len(prp.Key)
	if len(plainText)%k != 0 {
		plainText = pKCS7Pad(plainText, k)
	}
	block, err := aes.NewCipher(prp.Key)
	if err != nil {
		return nil, err
	}
	cipherData := make([]byte, len(plainText))
	blockMode := cipher.NewCBCEncrypter(block, prp.Iv)
	blockMode.CryptBlocks(cipherData, plainText)
	return cipherData, nil
}
