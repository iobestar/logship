package config

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

type Config struct {
	LogUnits []LogUnitConfig `yaml:"log_units"`
}

type LogUnitConfig struct {
	Id             string `yaml:"id"`
	FilePattern    string `yaml:"file_pattern"`
	LogPattern     string `yaml:"log_pattern"`
	DateTimeLayout string `yaml:"date_time_layout"`
}

var (
	defaultLogPattern     = "^(?P<datetime>\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}\\.\\d{3})\\s+(?P<level>\\w+).*"
	defaultDateTimeLayout = "2006-01-02 15:04:05.000"

	defaultLogUnit = LogUnitConfig{
		LogPattern:     defaultLogPattern,
		DateTimeLayout: defaultDateTimeLayout,
	}
)

func (lu *LogUnitConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {

	*lu = defaultLogUnit

	type plain LogUnitConfig
	if err := unmarshal((*plain)(lu)); err != nil {
		return err
	}
	return nil
}

func Parse(filename string) (*Config, error) {

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