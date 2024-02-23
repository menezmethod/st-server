package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port  string `mapstructure:"PORT"`
	DBUrl string `mapstructure:"DB_URL"`
}

func LoadConfig() (config Config, err error) {
	viper.AutomaticEnv()

	requiredVars := []string{"DB_URL", "PORT"}
	for _, v := range requiredVars {
		if err := viper.BindEnv(v); err != nil {
			return Config{}, err
		}
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
	}
	return
}
