package config

import "C"
import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var (
	Server   ServerConfig
	Redis    RedisConfig
	BlogInfo BlogInfoConfig
	Smtp     SmtpConfig
)

func init() {
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Panicf("error when reading yaml: %v", err)
	}
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Panicf("error when unmarshal yaml: %v", err)
	}

	Server = config.ServerConfig
	Redis = config.RedisConfig
	Smtp = config.SmtpConfig
	BlogInfo = config.BlogInfoConfig
}
