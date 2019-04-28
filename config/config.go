package config

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os"
)

type Config struct {
	LogReaders []LogReaderConfig `yaml:"log_readers"`
}

type LogReaderConfig struct {
	Id             string `yaml:"id"`
	LogPattern     string `yaml:"log_pattern"`
	DateTimeLayout string `yaml:"date_time_layout"`
}

var (
	defaultLogPattern     = "^(?P<datetime>\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}\\.\\d{3}).*"
	defaultDateTimeLayout = "2006-01-02 15:04:05.000"

	defaultLogReader = LogReaderConfig{
		LogPattern:     defaultLogPattern,
		DateTimeLayout: defaultDateTimeLayout,
	}

	defaultConfig = &Config{
		LogReaders: []LogReaderConfig{defaultLogReader},
	}
)

func (lr *LogReaderConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {

	*lr = defaultLogReader

	type plain LogReaderConfig
	if err := unmarshal((*plain)(lr)); err != nil {
		return err
	}
	return nil
}

func (c *Config) GetLogReaderConfig(logReaderConfigId string) LogReaderConfig {
	for _, lrc := range c.LogReaders {
		if lrc.Id == logReaderConfigId {
			return lrc
		}
	}
	return defaultLogReader
}

func ParseConfig(filename string) (*Config, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return defaultConfig, nil
	}

	cfgBytes, err := ioutil.ReadFile(filename)
	if nil != err {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(cfgBytes, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
