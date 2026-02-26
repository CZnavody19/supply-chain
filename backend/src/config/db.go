package config

type DBConfig struct {
	ConnectionURI string
	Username      string
	Password      string
}

func loadDbConfig() DBConfig {
	dbConfig := &DBConfig{}
	v := configViper("db")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	v.Unmarshal(dbConfig)
	return *dbConfig
}
