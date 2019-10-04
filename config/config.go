package config

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os"
)

type Config struct {
	LogUnits []LogUnit `yaml:"log_units"`
}

type LogUnit struct {
	Id   string `yaml:"id"`
	Glob string `yaml:"glob"`
}

var (
	defaultConfig = Config{
		LogUnits: []LogUnit{},
	}
)

func ParseConfig(filename string) (Config, error) {
	cfgBytes := []byte(os.Getenv("LOGSHIP_CONFIG"))
	if len(cfgBytes) == 0 {
		if len(filename) == 0 {
			return defaultConfig, nil
		}

		if _, err := os.Stat(filename); os.IsNotExist(err) {
			return defaultConfig, nil
		}

		var err error
		cfgBytes, err = ioutil.ReadFile(filename)
		if err != nil {
			return defaultConfig, err
		}

		if len(cfgBytes) == 0 {
			return defaultConfig, nil
		}
	}
	config := &Config{}
	if err := yaml.Unmarshal(cfgBytes, config); nil != err {
		return defaultConfig, err
	}
	return *config, nil
}
