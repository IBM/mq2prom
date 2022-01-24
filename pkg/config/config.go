package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"

	manager "github.ibm.com/chandergovind/mq2p/pkg/metric-manager"
)

type Config struct {
	MqConfig MqConfig `yaml:"mq"`
	Metrics  []manager.MetricConfig
}

type MqConfig struct {
	Server string
}

// ReadConfig reads the provided file for configuration
func ReadConfig(filename string) (Config, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error in reading config file: ")
		return Config{}, err
	}

	return ParseConfig(content)
}

// ParseConfig parses configuration byte array into configuration data struct
func ParseConfig(content []byte) (Config, error) {
	var config Config
	err := yaml.Unmarshal(content, &config)
	if err != nil {
		fmt.Printf("Error in unmarshalling config file.")
		return Config{}, err
	}

	return config, nil
}
