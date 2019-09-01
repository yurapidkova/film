package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Api struct{
		Key string
		Endpoint string
	}
	Run func()
}

var config Config
var initialized = false
var defaultPath = ".env"

func Read(path ...string)(*Config, error) {
	if initialized {
		return &config, nil
	}
	dirPath := defaultPath
	if len(path) > 0 {
		dirPath = path[0]
	}
	data, e := ioutil.ReadFile(dirPath)
	if e != nil {
		return &config, e
	}
	initialized = true
	return &config, yaml.Unmarshal(data, &config)
}


func Get() *Config  {
	if !initialized {
		_, _ = Read(defaultPath)
	}
	return &config
}