package config

import (
	"github.com/spf13/viper"
	"github.com/wncbb/be_blog/internal/pkg/env"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var isProduct bool

var configFilePath = "./configs/be_blog"
var configDevFileName = "dev_config.yaml"
var configProFileName = "pro_config.yaml"

var config Config

func Init() {
	env.IsProduct()
	viper.SetConfigType("yaml")
	configFile := os.Getenv("GOPATH") + "/src/github.com/wncbb/be_blog/configs/be_blog/"
	if env.IsProduct() {
		configFile = configFile + "/" + configProFileName
	} else {
		configFile = configFile + "/" + configDevFileName
	}
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	config.ip = viper.GetString("IP")
	config.port = viper.GetInt("Port")

	logrus.Debugf("read ip=%s, port=%d", config.ip, config.port)
}

func GetIP() string {
	return config.ip
}

func GetPort() int {
	return config.port
}

func GetRunAddr() string {
	return fmt.Sprintf("%s:%d", GetIP(), GetPort())
}
