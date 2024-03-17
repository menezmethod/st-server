package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Port          string `mapstructure:"PORT"`
	AuthSvcUrl    string `mapstructure:"AUTH_SVC_URL"`
	JournalSvcUrl string `mapstructure:"JOURNAL_SVC_URL"`
	ApiVersion    string `mapstructure:"API_VERSION"`
	JWTSecretKey  string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig() (config Config, err error) {
	viper.AutomaticEnv()
	requiredVars := []string{"AUTH_SVC_URL", "PORT", "JWT_SECRET_KEY", "JOURNAL_SVC_URL", "API_VERSION"}

	for _, v := range requiredVars {
		if err := viper.BindEnv(v); err != nil {
			return Config{}, err
		}
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Error unmarshaling configs: %v", err)
	}
	return
}
