package config

type Configuration interface {
	GetString(key string) (string, error)
	GetStringList(key string) ([]string, error)
	GetInt(key string) (int, error)
	GetBoolean(key string) (bool, error)
}

var cache = make(map[string]Configuration)

func Get(name, configType string) (config Configuration, err error) {
	config = cache[name]
	if config != nil {
		return config, nil
	}

	switch configType {
	case "mysql":
		{
			config, err = createMySQLConfiguration(name)
		}
	case "json":
		{
			config, err = createJsonConfiguration("config/" + name + ".json")
		}
	case "env":
		{
			config, err = createEnvConfiguration(name)
		}
	}

	cache[name] = config

	return
}

func generateKey(prefix, key string) string {
	return prefix + "_" + key
}
