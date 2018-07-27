package config

import (
	"flag"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
)

var conf config

type config struct {
	Db1 DbConfig    `yaml:"db1"`
	Db5 DbConfig    `yaml:"db5"`
}

type DbConfig struct {
	Type     string `yaml:"type"`
	Addr     string `yaml:"addr"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func init() {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.yaml", "config for application")
	flag.Parse()

	LoadConfig(configPath)
}

func LoadConfig(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
}

func GetConfig() *config {
	return &conf
}
