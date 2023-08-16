package properties_reader

import (
	properties_model "Golang-Gin-Gonic/properties/model"
	"fmt"

	"github.com/spf13/viper"
)

var Config properties_model.YamlConfig

func init() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../Golang-Gin-Gonic/")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config File is Missing")
		} else {
			fmt.Println("Another Error")
			fmt.Println(err)
		}
		panic(err)
	}

	viper.Unmarshal(&Config)
	fmt.Println("Env Variable is Read Successfully")
}
