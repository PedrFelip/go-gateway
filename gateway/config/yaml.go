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
	Server Port    `yaml:"server"`
	Routes []Route `yaml:"routes"`
}

type Port struct {
	Port int `yaml:"port"`
}

type Route struct {
	Path      string `yaml:"path"`
	Target    string `yaml:"target"`
	Protected bool   `yaml:"protected"`
	RateLimit int    `yaml:"rate_limit"`
}

func LoadConfig(path []byte) (*Config, error) {
	var config Config
	// Se a estrutura do YAML for inválida, retorna o erro.
	if err := yaml.Unmarshal(path, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
