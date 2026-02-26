package config

import "time"

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func loadServerConfig() ServerConfig {
	serverConfig := &ServerConfig{}
	v := configViper("server")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	v.Unmarshal(serverConfig)
	return *serverConfig
}
