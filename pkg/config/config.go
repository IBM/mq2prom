package config

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	"gopkg.in/yaml.v3"

	manager "github.com/IBM/mq2prom/pkg/metric-manager"
)

type Config struct {
	MqConfig MqConfig `yaml:"mq"`
	Metrics  []manager.MetricConfig
}

type MqConfig struct {
	Server string
	Topic  string
}

// ReadConfig reads the provided file for configuration
func ReadConfig(filename string) (Config, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error("Error in reading config file.")
		return Config{}, err
	}

	return ParseConfig(content)
}

// ParseConfig parses configuration byte array into configuration data struct
func ParseConfig(content []byte) (Config, error) {
	var config Config
	err := yaml.Unmarshal(content, &config)
	if err != nil {
		log.Error("Error in unmarshalling config file.")
		return Config{}, err
	}

	return config, nil
}
