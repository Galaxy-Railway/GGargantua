package common

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type ProjectConfig struct {
	Log    LogConfig    `yaml:"log"`
	Listen ListenConfig `yaml:"listen"`
}

type LogConfig struct {
	EnvType    EnvType  `yaml:"env_type"`
	OutputPath []string `yaml:"output_path"`
}

type ListenConfig struct {
	Network string `yaml:"network"`
	Address string `yaml:"address"`
}

type EnvType string

var (
	ProductEnv EnvType = "product"
	DevelopEnv EnvType = "develop"

	configPath = "./config.yaml"
)

func SetConfigPath(path string) {
	configPath = path
}

func NewConfig() (*ProjectConfig, error) {
	f, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	conf := &ProjectConfig{}
	err = yaml.Unmarshal(bytes, conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
