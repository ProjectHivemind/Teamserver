package conf

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// ConfOptions reads all conf data.
type ConfOptions struct {
	Database  map[string]string              `yaml:"database"`
	Restapi   map[string]string              `yaml:"restapi"`
	Users     []map[string]string            `yaml:"users"`
	Plugins   []map[string]map[string]string `yaml:"plugins"`
	Listeners []map[string]map[string]string `yaml:"listeners"`
}

// GetConf pulls the data from the config.yaml
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
