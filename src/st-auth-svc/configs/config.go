package configs

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port         string `mapstructure:"PORT"`
	DBUrl        string `mapstructure:"DB_URL"`
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	requiredVars := []string{"DB_URL", "PORT", "JWT_SECRET_KEY"}
	for _, v := range requiredVars {
		if err := viper.BindEnv(v); err != nil {
			return Config{}, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
