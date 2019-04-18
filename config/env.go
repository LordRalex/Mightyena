package config

import (
	"os"
	"strconv"
	"strings"
)

type envConfiguration struct {
	prefix string
}

func createEnvConfiguration(prefix string) (cfg Configuration, err error) {
	cfg = &envConfiguration{prefix: prefix}
	return
}

func (fc *envConfiguration) GetString(key string) (string, error) {
	val, _ := os.LookupEnv(generateKey(fc.prefix, strings.ToUpper(key)))
	return val, nil
}

func (fc *envConfiguration) GetStringList(key string) ([]string, error) {
	result, err := fc.GetString(key)

	if err != nil {
		return make([]string, 0), err
	}

	return strings.Split(result, "|"), nil
}

func (fc *envConfiguration) GetInt(key string) (int, error) {
	result, err := fc.GetString(key)

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(result)
}

func (fc *envConfiguration) GetBoolean(key string) (bool, error) {
	result, err := fc.GetString(key)

	if err != nil {
		return false, err
	}

	switch strings.ToLower(result) {
	case "true":
	case "1":
	case "y":
	case "yes":
		return true, nil
	}

	return false, nil
}
