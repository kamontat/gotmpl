package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	underlay map[interface{}]interface{}
}

func (c *Config) Variable() map[interface{}]interface{} {
	return c.underlay
}

func New(filetype string, filepath string) *Config {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalln(err)
	}

	if filetype == "yaml" || filetype == "yml" {
		var underlay = make(map[interface{}]interface{})

		err = yaml.Unmarshal(content, &underlay)
		if err != nil {
			log.Fatalln(err)
		}

		return &Config{
			underlay: underlay,
		}
	}

	log.Fatalln("invalid config type")
	return nil
}
