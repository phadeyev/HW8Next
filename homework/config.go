package main

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Logger LoggerConfig `yaml:"logger"`
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Addr    string `yaml:"addr"`
	MongoDB string `yaml:"mongodb"`
}

type LoggerConfig struct {
	DebugLevel string `yaml:"debuglevel"`
}

func (c *Config) ReadConfig(path string) error { //3
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	data, _ := ioutil.ReadAll(f)

	err = yaml.Unmarshal(data, &c)

	if err != nil {
		return err
	}

	return nil
}
