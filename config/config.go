package config

import "github.com/spf13/viper"

type Env struct {
	DBUser string
	DBPass string
	DBName string
	DBHost string
	DBPort int
	ServerPort string
	JWTSecret string
}

var ENV *Env

func LoadConfig() (*Env, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	env := &Env{
		DBUser: viper.GetString("DB_USER"),
		DBPass: viper.GetString("DB_PASS"),
		DBName: viper.GetString("DB_NAME"),
		DBHost: viper.GetString("DB_HOST"),
		DBPort: viper.GetInt("DB_PORT"),
		ServerPort: viper.GetString("SERVER_PORT"),
		JWTSecret: viper.GetString("JWT_SECRET"),
	}

	ENV = env
	return env, nil
}