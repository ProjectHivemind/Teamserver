package conf

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type ConfOptions struct {
	Database  map[string]string              `yaml:"database"`
	Restapi   map[string]string              `yaml:"restapi"`
	Listeners []map[string]map[string]string `yaml:"listeners"`
}

func (c *ConfOptions) GetConf(path string) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}