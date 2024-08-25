package jsonenv

import (
	"os"
	"strings"

	"github.com/sinemah/jsonenv/encrypt/aescrypt"
)

type Crypter struct {
	aesCrypter *aescrypt.Crypter
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

func loadAesCrypter() *aescrypt.Crypter {
	c, err := aescrypt.NewCrypter(os.Getenv("JSONENV_AES_FILE"))

	if err != nil {
		return nil
	}

	return c
}
