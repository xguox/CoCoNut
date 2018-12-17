package config

import (
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	Server   ServerConf
	Database DatabaseConf
}

type ServerConf struct {
	Port          string
	SecretBaseKEY string
}

type DatabaseConf struct {
	Host     string
	Port     int
	Basename string
	Pwd      string
}

var Conf Configuration

func init() {
	viper.SetConfigName("conf")   // no need to include file extension
	viper.AddConfigPath("config") // set the path of your config file

	if err := viper.ReadInConfig(); err != nil {
		log.Println("Config file not found...")
	}
	err := viper.Unmarshal(&Conf)

	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}
