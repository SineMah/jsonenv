package aescrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"os"
)

type Crypter struct {
	key []byte
}

type EncryptValue struct {
	Iv    string `json:"iv"`
	Value string `json:"value"`
}

func NewCrypter(file string) (*Crypter, error) {
	key, err := os.ReadFile(file)

	if err != nil {
		return nil, err
	}

	return &Crypter{
		key: key,
	}, nil
}

func (c Crypter) EncryptValue(value string) (string, error) {
	block, err := aes.NewCipher([]byte(c.key))

	if err != nil {
		return "", err
	}

	iv, err := GenerateIv()

	if err != nil {
		return "", err
	}

	length := len(value)
	textBlock := make([]byte, length)

	if length%16 != 0 {
		extendBlock := 16 - (length % 16)
		textBlock = make([]byte, length+extendBlock)
		copy(textBlock[length:], bytes.Repeat([]byte{uint8(extendBlock)}, extendBlock))
	}

	copy(textBlock, value)

	ciphertext := make([]byte, len(textBlock))

	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, textBlock)

	b, err := json.Marshal(EncryptValue{
		Iv:    iv,
		Value: base64.StdEncoding.EncodeToString(ciphertext),
	})

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

func (c Crypter) DecryptValue(value string) (string, error) {
	j, err := base64.StdEncoding.DecodeString(value)

	if err != nil {
		return "", err
	}

	var ev EncryptValue
	err = json.Unmarshal(j, &ev)

	if err != nil {
		return "", err
	}

	v, err := base64.StdEncoding.DecodeString(ev.Value)

	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(c.key)

	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, []byte(ev.Iv))
	mode.CryptBlocks(v, v)

	return string(PKCS5UnPadding(v)), nil

}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])

	return src[:(length - unpadding)]
}
