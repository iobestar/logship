package config

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
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
	defaultLogPattern     = "^(?P<datetime>\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}\\.\\d{3})\\s+(?P<level>\\w+).*"
	defaultDateTimeLayout = "2006-01-02 15:04:05.000"

	defaultLogReader = LogReaderConfig{
		Id:"default",
		LogPattern:     defaultLogPattern,
		DateTimeLayout: defaultDateTimeLayout,
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

	if len(c.LogReaders) == 0 {
		return defaultLogReader
	}

	for _, lrc := range c.LogReaders {
		if lrc.Id == logReaderConfigId {
			return lrc
		}
	}
	return defaultLogReader
}

func ParseConfig(filename string) (*Config, error) {

	bytes, err := getBytes(filename)
	if nil != err {
		return nil, err
	}
	return parseBytes(bytes)
}

func parseBytes(bytes []byte) (*Config, error) {

	cfg := &Config{}
	if err := yaml.Unmarshal(bytes, cfg); nil != err {
		panic(err)
	}

	var config Config
	err := yaml.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func getBytes(filename string) ([]byte, error) {

	cfgBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return cfgBytes, nil
}
