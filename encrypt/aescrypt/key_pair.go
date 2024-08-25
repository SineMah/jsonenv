package aescrypt

import (
	"bufio"
	"crypto/rand"
	"math/big"
	"os"
)

type Crypt struct {
	Key string
	Iv  string
}

type Config struct {
	File    string
	KeyType string
}

func NewKey(file string) *Config {
	return &Config{
		File:    file,
		KeyType: "aes-256",
	}
}

func (e Config) ReadKey() (string, string, error) {
	f, err := os.Open(e.File)

	if err != nil {
		return "", "", err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var rows []string

	for scanner.Scan() {
		row := scanner.Text()

		if len(row) > 0 {
			rows = append(rows, row)
		}
	}

	if len(rows) != 2 {
		return "", "", nil
	}

	return rows[0], rows[1], nil
}

func (e Config) DeleteKey() error {
	return os.Remove(e.File)
}

func (e Config) GenerateKey() error {
	key, err := randomSequence(32)

	if err != nil {
		return err
	}

	_, err = os.Stat(e.File)

	if err == nil {
		e.DeleteKey()
	}

	return os.WriteFile(e.File, []byte(key), 0644)
}

func GenerateIv() (string, error) {
	iv, err := randomSequence(16)

	if err != nil {
		return "", err
	}

	return iv, nil
}

func randomSequence(n int) (string, error) {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	bi := big.NewInt(int64(len(letters)))
	b := make([]rune, n)

	for i := range b {
		ra, err := rand.Int(rand.Reader, bi)

		if err != nil {
			return "", err
		}

		b[i] = letters[ra.Int64()]
	}

	return string(b), nil
}
