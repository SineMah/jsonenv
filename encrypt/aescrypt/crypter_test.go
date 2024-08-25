package aescrypt

import (
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAesCrypter(t *testing.T) {
	t.Run("test aes encrypt value", func(t *testing.T) {
		e := NewKey(".crypt")

		e.GenerateKey()

		c, err := NewCrypter(".crypt")

		assert.Nil(t, err)

		s, err := c.EncryptValue("test")

		assert.Nil(t, err)

		data, err := base64.StdEncoding.DecodeString(s)

		assert.Nil(t, err)

		var ev EncryptValue
		err = json.Unmarshal([]byte(data), &ev)

		assert.Nil(t, err)

		origin, err := base64.StdEncoding.DecodeString(ev.Value)

		assert.Nil(t, err)

		assert.NotEqual(t, "test", string(origin))
	})

	t.Run("test aes decrypt value", func(t *testing.T) {
		c, err := NewCrypter(".crypt")

		assert.Nil(t, err)

		s, err := c.EncryptValue("test")

		assert.Nil(t, err)

		v, err := c.DecryptValue(s)

		assert.Nil(t, err)

		assert.Equal(t, "test", v)
	})
}
