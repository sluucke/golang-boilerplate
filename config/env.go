package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	ServerAddr string `mapstructure:"SERVER_ADDRESS"`
	AppEnv     string `mapstructure:"ENVIROMENT"`
	Port       string `mapstructure:"PORT"`
}

func NewEnv() *Env {
	env := Env{}

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("Cant find enviroment file", err)
	}

	err = viper.Unmarshal(&env)

	if err != nil {
		log.Fatal("Cant read env file", err)
	}

	return &env
}
