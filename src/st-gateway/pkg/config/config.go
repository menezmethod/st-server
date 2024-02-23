package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port           string `mapstructure:"PORT"`
	AuthSvcUrl     string `mapstructure:"AUTH_SVC_URL"`
	JournalSvcUrl  string `mapstructure:"JOURNAL_SVC_URL"`
	ApiVersion     string `mapstructure:"API_VERSION"`
	JWTSecretKey   string `mapstructure:"JWT_SECRET_KEY"`
	AllowedOrigins []string
}

func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()

	requiredVars := []string{"AUTH_SVC_URL", "PORT", "JWT_SECRET_KEY", "JOURNAL_SVC_URL", "API_VERSION"}
	for _, v := range requiredVars {
		if err := viper.BindEnv(v); err != nil {
			log.Fatalf("Failed to bind environment variable: %s, err: %v", v, err)
			return nil, err
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
		return nil, err
	}

	config.AllowedOrigins = []string{"*"}

	return &config, nil
}
