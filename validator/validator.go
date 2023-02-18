package validator

import (
	"fmt"
	"os"
)

func Config(name string, config map[string]*string) string {
	// Check config
	yamlFile := "oapi-cfg.yaml"
	if _, err := os.Stat(yamlFile); os.IsNotExist(err) {
		return fmt.Sprintf("%s does not exist", yamlFile)
	}

	// Check spec and spec config exists
	yamlFile = fmt.Sprintf("%s/server.cfg.yaml", *config["spec-dir"])
	if _, err := os.Stat(yamlFile); os.IsNotExist(err) {
		return fmt.Sprintf("spec server config (%s) does not exist", yamlFile)
	}
	yamlFile = fmt.Sprintf("%s/%s.yaml", *config["spec-dir"], name)
	if _, err := os.Stat(yamlFile); os.IsNotExist(err) {
		return fmt.Sprintf("spec file (%s) does not exist", yamlFile)
	}
	return ""
}
