package config

type LoggingConfig struct {
	EnableDebugLogger bool
	EnableFileLogger  bool
	FileLogLevel      string
	FileLogOutput     string
	FileMaxSize       int
	FileMaxBackups    int
	FileMaxAge        int
}

func loadLoggingConfig() LoggingConfig {
	loggingConfig := &LoggingConfig{}
	v := configViper("logging")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	v.Unmarshal(loggingConfig)
	return *loggingConfig
}
