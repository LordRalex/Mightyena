package config

type Configuration interface {
	GetString(key string) (string, error)
	GetStringList(key string) ([]string, error)
	GetStringMap(key string) (map[string]string, error)
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
			config, err = createMySQLConfiguration()
		}
	case "json":
		{
			config, err = createJsonConfiguration("config/" + name + ".json")
		}
	}

	cache[name] = config

	return
}
