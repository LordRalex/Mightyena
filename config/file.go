package config

import (
	"encoding/json"
	"os"
	"strings"
)

type jsonConfiguration struct {
	mapping interface{}
}

func createJsonConfiguration (path string) (config Configuration, err error) {
	cfg := &jsonConfiguration{}
	err = cfg.load(path)
	return cfg, err
}

func (fc *jsonConfiguration) load(path string) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(fc.mapping)
}

func (fc *jsonConfiguration) GetString(key string) (string, error) {
	data := get(fc.mapping, key)

	str, ok := data.(string)
	if !ok {
		str = ""
	}

	return str, nil
}
func (fc *jsonConfiguration) GetStringList(key string) ([]string, error) {
	data := get(fc.mapping, key)

	str, ok := data.([]string)
	if !ok {
		str = make([]string, 0)
	}

	return str, nil
}

func (fc *jsonConfiguration) GetStringMap(key string) (map[string]string, error) {
	data := get(fc.mapping, key)

	str, ok := data.(map[string]string)
	if !ok {
		str = make(map[string]string)
	}

	return str, nil
}

func (fc *jsonConfiguration) GetInt(key string) (int, error) {
	data := get(fc.mapping, key)

	str, ok := data.(int)
	if !ok {
		str = 0
	}

	return str, nil
}

func (fc *jsonConfiguration) GetBoolean(key string) (bool, error) {
	data := get(fc.mapping, key)

	str, ok := data.(bool)
	if !ok {
		str = false
	}

	return str, nil
}

func get(data interface{}, key string) interface{} {
	keyParts := strings.SplitN(key, ".", 2)

	mapping, ok := data.(map[string]interface{})
	if !ok {

		newMap, ok := data.(map[string]string)
		if !ok {
			return nil
		}

		mapping = make(map[string]interface{})

		for k, v := range newMap {
			mapping[k] = v
		}
	}

	if len(keyParts) == 1 {
		return mapping[keyParts[0]]
	} else {
		return get(mapping[keyParts[0]], keyParts[1])
	}
}