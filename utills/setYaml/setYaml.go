package setYaml

import (
	"autoSql-cobra/config"
	"autoSql-cobra/utills/getRunDir"
	"fmt"
	"github.com/spf13/viper"
)

func SetYaml[T string | []string](key string, value T) (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetConfigFile(fmt.Sprintf("%s/%s.yaml", getRunDir.RunDir(), "config"))

	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println(err.Error())
		} else {
			fmt.Println(err.Error())
		}
	}

	viper.Set(key, value)
	err = viper.WriteConfigAs(fmt.Sprintf("%s/%s.yaml", getRunDir.RunDir(), "config"))

	if err = viper.Unmarshal(&config.Config); err != nil {
		fmt.Printf("Serialization configuration failed.%s\n", err)
		return
	}

	return err
}
