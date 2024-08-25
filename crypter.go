package jsonenv

import (
	"os"
	"strings"

	"github.com/sinemah/jsonenv/encrypt/aes_crypt"
)

type Crypter struct {
	aesCrypter *aes_crypt.Crypter
}

func LoadCrypter() *Crypter {
	return &Crypter{
		aesCrypter: loadAesCrypter(),
	}
}

func (c Crypter) Decrypt(value string) string {
	var err error
	var encryptedValue string

	if strings.HasPrefix(value, "encrypted:") == false {
		return value
	}

	v := strings.Split(value, ":")

	if len(v) < 3 {
		return ""
	}

	switch true {
	case v[1] == "aes" && c.aesCrypter != nil:
		encryptedValue, err = c.aesCrypter.DecryptValue(v[2])
	default:
		return ""
	}

	if err != nil {
		return ""
	}

	return encryptedValue
}

func loadAesCrypter() *aes_crypt.Crypter {
	c, err := aes_crypt.NewCrypter(os.Getenv("JSONENV_AES_FILE"))

	if err != nil {
		return nil
	}

	return c
}
