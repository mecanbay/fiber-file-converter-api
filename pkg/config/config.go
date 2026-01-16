package config

import (
	"time"

	"github.com/spf13/viper"
)

var (
	ErrConfigFileNotFound = "Config file not found or not readable"
	ErrUnmarshallConfig   = "Config decode error"
)

type AppConfig struct {
	AppName                 string        `mapstructure:"app_name"`
	Hostname                string        `mapstructure:"hostname"`
	Port                    string        `mapstructure:"port"`
	Env                     string        `mapstructure:"env"`
	LogPath                 string        `mapstructure:"log_path"`
	LogFile                 string        `mapstructure:"log_file"`
	ReadTimeout             time.Duration `mapstructure:"read_timeout"`
	WriteTimeout            time.Duration `mapstructure:"write_timeout"`
	IdleTimeout             time.Duration `mapstructure:"idle_timeout"`
	GracefulShutdownTimeout time.Duration `mapstructure:"graceful_shutdown_timeout"`
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
