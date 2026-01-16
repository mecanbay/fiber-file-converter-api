package config

import "github.com/spf13/viper"

var (
	ErrConfigFileNotFound = "Config file not found or not readable"
	ErrUnmarshallConfig   = "Config decode error"
)

type AppConfig struct {
	AppName  string `mapstructure:"app_name"`
	Hostname string `mapstructure:"hostname"`
	Port     string `mapstructure:"port"`
}

type Config struct {
	App AppConfig `mapstructure:"app"`
}

func Load() *Config {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("../config")

	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var cfg Config
	err = v.Unmarshal(&cfg)
	if err != nil {
		panic(ErrUnmarshallConfig)
	}

	return &cfg
}
