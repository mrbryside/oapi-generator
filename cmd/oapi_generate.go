package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Initialize your application",
	Long:  `Initialize your application by creating a new configuration file and directory structure`,
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		yamlFile := "oapi-cfg.yaml"
		if _, err := os.Stat(yamlFile); os.IsNotExist(err) {
			return fmt.Errorf("%s does not exist", yamlFile)
		}

		// Read the contents of the YAML file into a byte slice
		data, err := ioutil.ReadFile("oapi-cfg.yaml")
		if err != nil {
			log.Fatalf("failed to read YAML file: %v", err)
		}

		// Parse the YAML data into a map[string]interface{} object
		var config map[string]interface{}
		err = yaml.Unmarshal(data, &config)
		if err != nil {
			log.Fatalf("failed to parse YAML data: %v", err)
		}

		yamlFile = fmt.Sprintf("%s/server.cfg.yaml", config["spec-dir"])
		if _, err := os.Stat(yamlFile); os.IsNotExist(err) {
			return fmt.Errorf("spec server config (%s) does not exist", yamlFile)
		}
		yamlFile = fmt.Sprintf("%s/%s.yaml", config["spec-dir"], name)
		if _, err := os.Stat(yamlFile); os.IsNotExist(err) {
			return fmt.Errorf("spec file (%s) does not exist", yamlFile)
		}

		// Run the command "go run main.go init"
		err = createFolder()
		if err != nil {
			return fmt.Errorf("error running init folder: %v", err)
		}

		genCmd := exec.Command("make", "generate-server", fmt.Sprintf("name=%s", name), fmt.Sprintf("genPath=%s", config["gen-dir"]), fmt.Sprintf("specPath=%s", config["spec-dir"]))
		genCmd.Dir = "./"
		_, err = genCmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("error running generate: %v", err)
		}
		fmt.Println("generated")
		return nil
	}}

func init() {
	generateCmd.Flags().String("name", "", "Name of the application")
	rootCmd.AddCommand(generateCmd)
}
