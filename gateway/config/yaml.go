package config

import (
	"gopkg.in/yaml.v3"
)

// Estrutura de configuração do gateway
/*
 server:
  port: 1234

routes:
  - path:
    target:
    protected: false
    rate_limit:
*/
type Config struct {
	Server ServerConfig    `yaml:"server"`
	Routes []Route `yaml:"routes"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type Route struct {
	Path      string `yaml:"path"`
	Target    string `yaml:"target"`
	Protected bool   `yaml:"protected"`
	RateLimit int    `yaml:"rate_limit"`
}

func LoadConfig(yamlBytes []byte) (*Config, error) {
	var config Config
	// Se a estrutura do YAML for inválida, retorna o erro.
	if err := yaml.Unmarshal(yamlBytes, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
