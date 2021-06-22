package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	FeaturesPaths    []string    `mapstructure:"features_path"`
	ReadinessTimeout string      `mapstructure:"readiness_timeout"`
	StopOnFailure    bool        `mapstructure:"stop_on_failure"`
	Resources        []*Resource `mapstructure:"resources"`
}

type Resource struct {
	Name    string            `mapstructure:"name"`
	Type    string            `mapstructure:"type"`
	Options map[string]string `mapstructure:"options"`
}

func ReadConfig(path string) (cfg *Config, err error) {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(path)
	err = v.ReadInConfig()
	if err != nil {
		err = fmt.Errorf("Fatal error config file: %w \n", err)
		return
	}

	err = v.Unmarshal(&cfg)
	return
}
