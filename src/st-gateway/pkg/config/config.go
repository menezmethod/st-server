package config

import "github.com/spf13/viper"

type Config struct {
	Port          string `mapstructure:"PORT"`
	AuthSvcUrl    string `mapstructure:"AUTH_SVC_URL"`
	JournalSvcUrl string `mapstructure:"JOURNAL_SVC_URL"`
	ApiVersion    string `mapstructure:"API_VERSION"`
	JWTSecretKey  string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./../config/envs")
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
