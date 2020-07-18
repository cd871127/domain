package config

import "github.com/spf13/viper"

func Load() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName("config.yaml")
	v.AddConfigPath("config/")
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {

		return nil, err

	}
	return v, nil
}
