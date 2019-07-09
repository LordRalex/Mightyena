package config

import (
	"github.com/lordralex/mightyena/database"
	"strconv"
	"strings"
)

type mySQLConfiguration struct {
	prefix string
}

func createMySQLConfiguration(prefix string) (cfg Configuration, err error) {
	cfg = &mySQLConfiguration{prefix: prefix}
	return
}

func (fc *mySQLConfiguration) GetString(key string) (string, error) {
	db := database.GetConnection()
	setting := &dbSetting{}

	res := db.Table("settings").Where("key = ?", generateKey(fc.prefix, key)).FirstOrInit(&setting)

	return setting.Value, res.Error
}

func (fc *mySQLConfiguration) GetStringList(key string) ([]string, error) {
	result, err := fc.GetString(key)

	if err != nil {
		return make([]string, 0), err
	}

	return strings.Split(result, "|"), nil
}

func (fc *mySQLConfiguration) GetInt(key string) (int, error) {
	result, err := fc.GetString(key)

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(result)
}

func (fc *mySQLConfiguration) GetBoolean(key string) (bool, error) {
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

type dbSetting struct {
	Key   string
	Value string
}
