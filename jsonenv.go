package jsonenv

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"strconv"
)

func Load(filenames ...string) (err error) {
	filenames = filenamesOrDefault(filenames)

	for _, filename := range filenames {
		err = loadFile(filename)

		if err == nil {
			break
		}
	}

	return err
}

func filenamesOrDefault(filenames []string) []string {

	if len(filenames) == 0 {

		return []string{"env.json"}
	}

	return filenames
}

func loadFile(filename string) error {
	file, err := os.Open(filename)

	defer file.Close()

	if err != nil {
		log.Panic(err)
	}

	bytes, err := io.ReadAll(file)

	if err != nil {
		return err
	}

	envMap, err := Unmarshal(bytes)

	for entry := range envMap {
		_ = os.Setenv(entry, envMap[entry])
	}

	return err
}

func Unmarshal(bytes []byte) (map[string]string, error) {
	envRaw := make(map[string]any)

	err := json.Unmarshal(bytes, &envRaw)

	if err != nil {
		return nil, err
	}

	if len(envRaw) == 0 {
		return nil, errors.New("empty env file loaded")
	}

	m := flattenMap(envRaw, "")

	return m, nil
}

func flattenMap(input map[string]any, prefix string) map[string]string {
	result := make(map[string]string)

	for key, value := range input {
		fullKey := key
		if prefix != "" {
			fullKey = prefix + "." + key
		}

		if reflect.ValueOf(value).Kind() == reflect.Map {
			nested := flattenMap(value.(map[string]any), fullKey)

			for k, v := range nested {
				result[k] = v
			}
		} else {
			result[fullKey] = ConvertAnyToString(value)
		}
	}

	if prefix == "" {
		for key, value := range result {
			if key[0] == '.' {
				delete(result, key)
				result[key[1:]] = value
			}
		}
	}

	return result
}

func ConvertAnyToString(value any) string {
	switch v := value.(type) {
	case string:
		return v
	case int, int32, int64, float32, float64:
		return fmt.Sprintf("%v", v)
	case bool:
		return strconv.FormatBool(v)
	default:
		return ""
	}
}
