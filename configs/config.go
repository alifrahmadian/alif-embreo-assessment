package config

type Config struct {
	DB  DBConfig
	Env string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}
