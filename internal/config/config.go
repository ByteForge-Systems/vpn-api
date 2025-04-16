package config

type Config struct {
	Port string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() *Config {
	LoadEnv()

	cfg := &Config{
		Port:       GetEnv("PORT"),
		DBHost:     GetEnv("DB_HOST"),
		DBPort:     GetEnv("DB_PORT"),
		DBUser:     GetEnv("DB_USER"),
		DBPassword: GetEnv("DB_PASSWORD"),
		DBName:     GetEnv("DB_NAME"),
	}

	if cfg.Port == "" {
		cfg.Port = "8080"
	}
	return cfg
}
