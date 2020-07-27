package common

import "github.com/spf13/viper"

func Load(configPath string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigName("config.yaml")
	v.AddConfigPath(configPath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil
}
