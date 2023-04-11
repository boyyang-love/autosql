package config

import (
	"autoSql-cobra/utills/getRunDir"
	"fmt"
	"github.com/spf13/viper"
)

type SqlConfig struct {
	Schema string   `mapstructure:"schema"`
	Fields []string `mapstructure:"fields"`
	Exts   []string `mapstructure:"exts"`
}

var Config SqlConfig

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetConfigFile(fmt.Sprintf("%s/%s.yaml", getRunDir.RunDir(), "config"))

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println(err.Error())
		} else {
			fmt.Println(err.Error())
		}
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Printf("Serialization configuration failed.%s\n", err)
		return
	}
}
