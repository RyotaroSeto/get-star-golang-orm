package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	GithubToken string `mapstructure:"GITHUB_TOKEN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// viper.AutomaticEnv()
	viper.BindEnv("GithubToken")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
