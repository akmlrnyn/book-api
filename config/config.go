package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	PORT int
	DB_USER string
	DB_PASSWORD string
	DB_DATABASE string
	DB_PORT int
	DB_HOST string
}

var ENV Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()

	if err != nil{
		log.Fatal(err.Error())
		return
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		log.Fatal(err.Error())
		return
	}

	fmt.Println("Load server success")

}