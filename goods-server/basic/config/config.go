package config

type AppConfig struct {
	Mysql struct {
		User     string
		Password string
		Host     string
		Port     int
		Database string
	}
}

var Con AppConfig
