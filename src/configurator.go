package main

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type ConfigStruct struct {
	DataBase struct {
		Name string
		Host string
		Port int
		User string
		Pass string
	}

	WebServer struct {
		Host        string
		Port        int
		BindingPath string
	}
}

func LoadConfig(configPath string) (*ConfigStruct, error) {
	var config ConfigStruct

	fConfig, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("can't open YAML config %q: %s", configPath, err)
	}

	fData, err := ioutil.ReadAll(fConfig)
	if err != nil {
		return nil, fmt.Errorf("can't read YAML file %q: %s", configPath, err)
	}

	err = yaml.Unmarshal(fData, &config)
	if err != nil {
		return nil, fmt.Errorf("can't unmarshal YAML file %q: %s", configPath, err)
	}

	return &config, nil
}
