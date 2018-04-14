package config

import (
	"io/ioutil"
	"os"

	"github.com/canhlinh/go-api/src/models"
	"gopkg.in/yaml.v2"
)

var config *models.Config

func Load(yamlPath string) {
	configFile, err := os.Open(yamlPath)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(configFile)
	if err != nil {
		panic(err)
	}

	config = &models.Config{}
	if err := yaml.Unmarshal(data, config); err != nil {
		panic(err)
	}
}

func Config() models.Config {
	return *config
}
