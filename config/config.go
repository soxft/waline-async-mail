package config

import "C"
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

var (
	templatesPath string
	configPath    string
)

var (
	OwnerTemplate string
	GuestTemplate string
)

func init() {
	testing.Init()
	flag.StringVar(&templatesPath, "t", "./templates", "模板文件路径")
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

	// read templates
	if _, err := ioutil.ReadDir(templatesPath); err != nil {
		log.Fatalf("templates dir not exists: %v", err)
	}
	// get file content
	_ownerTemplate, err := ioutil.ReadFile(templatesPath + "/owner.html")
	if err != nil {
		log.Fatalf("error when reading owner template: %v", err)
	}
	_guestTemplate, err := ioutil.ReadFile(templatesPath + "/guest.html")
	if err != nil {
		log.Fatalf("error when reading guest template: %v", err)
	}
	// set global variable
	OwnerTemplate = string(_ownerTemplate)
	GuestTemplate = string(_guestTemplate)
}
