package jsonenv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {

	t.Run("test valid environment config", func(t *testing.T) {
		bytes := []byte(`{"port": 1111, "port_as_string":"1110", "foo": {"bar": "baz"}, "is_true": true}`)
		r, err := Unmarshal(bytes)

		if err != nil {
			t.Errorf("Error unmarshalling: %v", err)
		}

		assert.Equal(t, "1111", r["port"])
		assert.Equal(t, "1110", r["port_as_string"])
		assert.Equal(t, "baz", r["foo.bar"])
		assert.Equal(t, "true", r["is_true"])
	})

	t.Run("test invalid environment config", func(t *testing.T) {
		bytes := []byte(`{"port: 1111}`)
		r, err := Unmarshal(bytes)

		assert.Error(t, err, "empty env file loaded")
		assert.Equal(t, 0, len(r))
	})
}

func TestConvertAnyToString(t *testing.T) {

	t.Run("test convert int to string", func(t *testing.T) {
		assert.Equal(t, "1", ConvertAnyToString(1))
		assert.Equal(t, "753", ConvertAnyToString(753))
	})

	t.Run("test convert float to string", func(t *testing.T) {
		assert.Equal(t, "1", ConvertAnyToString(1.00))
		assert.Equal(t, "1.001", ConvertAnyToString(1.001))
		assert.Equal(t, "11.11", ConvertAnyToString(11.11))
		assert.Equal(t, "-11.11", ConvertAnyToString(-11.11))
	})

	t.Run("test convert bool to string", func(t *testing.T) {
		assert.Equal(t, "true", ConvertAnyToString(true))
		assert.Equal(t, "false", ConvertAnyToString(false))
	})

	t.Run("test string is not modified", func(t *testing.T) {
		assert.Equal(t, "A smile can be more powerful than a roar. ~Cheshire Cat", ConvertAnyToString("A smile can be more powerful than a roar. ~Cheshire Cat"))
		assert.Equal(t, "Curiosity is the key to Wonderland. ~Cheshire Cat", ConvertAnyToString("Curiosity is the key to Wonderland. ~Cheshire Cat"))
	})

	t.Run("test convert anything to string", func(t *testing.T) {
		m := make(map[string]float64)

		m["e"] = 2.71828

		assert.Equal(t, "", ConvertAnyToString(m))
	})
}
