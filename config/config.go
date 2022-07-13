package config

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"testing"
)

var (
	Server   ServerConfig
	Redis    RedisConfig
	BlogInfo BlogInfoConfig
	Smtp     SmtpConfig
)

var configPath string

func init() {
	testing.Init()
	flag.StringVar(&configPath, "c", "./config.yaml", "配置文件路径")
	flag.Parse()

	// read config
	data, err := ioutil.ReadFile(configPath)
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
