package logging

type Logger interface {
	//Log(level Level, format string, args ...interface{})

	Info(format string, args ...interface{})

	Warning(format string, args ...interface{})

	Error(format string, args ...interface{})

	Critical(format string, args ...interface{})

	Debug(format string, args ...interface{})
}

var cache = make(map[string]Logger)

func GetLogger(name string) Logger {

	logger := cache[name]
	if logger == nil {
		logger = &prefixLogger{Name: name}
		cache[name] = logger
	}

	return logger
}
