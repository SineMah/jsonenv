package aescrypt

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptedEnv(t *testing.T) {
	t.Run("test generate key pair", func(t *testing.T) {
		keyFile := ".crypt"

		e := NewKey(keyFile)

		e.GenerateKey()
		e.GenerateKey()

		b, err := os.ReadFile(keyFile)

		assert.Nil(t, err)

		assert.Equal(t, 32, len(b))
	})
}
