package reader

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func ReadConfig() map[string]*string {
	// Parse the YAML data into a map[string]interface{} object
	// Read the contents of the YAML file into a byte slice
	var config map[string]*string
	data, err := ioutil.ReadFile("oapi-cfg.yaml")
	if err != nil {
		fmt.Errorf("failed to read YAML file: %v", err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Errorf("failed to parse YAML data: %v", err)
	}
	if config["spec-dir"] == nil || config["gen-dir"] == nil {
		fmt.Errorf("error oapi-cfg.yaml is wrong")
	}
	return config
}
