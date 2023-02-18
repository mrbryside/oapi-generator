package cmd

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

func createFolder() error {
	data, err := ioutil.ReadFile("oapi-cfg.yaml")
	if err != nil {
		log.Fatalf("failed to read YAML file: %v", err)
		return err
	}

	// Parse the YAML data into a map[string]interface{} object
	var config map[string]interface{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("failed to parse YAML data: %v", err)
		return err
	}

	// create folder in generated path
	err = os.MkdirAll(fmt.Sprintf("./%s/oapi", config["gen-dir"]), 0755)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
