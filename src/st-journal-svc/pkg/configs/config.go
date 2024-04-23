package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AuthSvcUrl   string `mapstructure:"AUTH_SVC_URL"`
	HelperSvcUrl string `mapstructure:"HELPER_SVC_URL"`
	DBUrl        string `mapstructure:"DB_URL"`
	Port         string `mapstructure:"PORT"`
}

func LoadConfig() (config Config, err error) {
	viper.AutomaticEnv()

	requiredVars := []string{"DB_URL", "PORT", "AUTH_SVC_URL"}
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
