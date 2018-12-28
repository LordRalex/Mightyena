package config

type mySQLConfiguration struct {
}

func createMySQLConfiguration() (cfg Configuration, err error) {
	cfg = &mySQLConfiguration{}
	return
}

func (fc *mySQLConfiguration) GetString(key string) (string, error) {
	return "", nil
}
func (fc *mySQLConfiguration) GetStringList(key string) ([]string, error) {
	return make([]string, 0), nil
}

func (fc *mySQLConfiguration) GetStringMap(key string) (map[string]string, error) {
	return make(map[string]string), nil
}

func (fc *mySQLConfiguration) GetInt(key string) (int, error) {
	return 0, nil
}

func (fc *mySQLConfiguration) GetBoolean(key string) (bool, error) {
	return false, nil
}
