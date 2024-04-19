package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ApiVersion    string `mapstructure:"API_VERSION"`
	AuthSvcUrl    string `mapstructure:"AUTH_SVC_URL"`
	HelperSvcUrl  string `mapstructure:"HELPER_SVC_URL"`
	JWTSecretKey  string `mapstructure:"JWT_SECRET_KEY"`
	JournalSvcUrl string `mapstructure:"JOURNAL_SVC_URL"`
	Port          string `mapstructure:"PORT"`
}

func LoadConfig() (config Config, err error) {
	viper.AutomaticEnv()
	requiredVars := []string{"AUTH_SVC_URL", "PORT", "JWT_SECRET_KEY", "JOURNAL_SVC_URL", "HELPER_SVC_URL", "API_VERSION"}

	for _, v := range requiredVars {
		if errBind := viper.BindEnv(v); errBind != nil {
			return Config{}, errBind
		}
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Error unmarshaling configs: %v", err)
	}
	return
}
