package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize your application",
	Long:  `Initialize your application by creating a new configuration file and directory structure`,
	Run: func(cmd *cobra.Command, args []string) {
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

		// create folder in generated path
		err = os.MkdirAll(fmt.Sprintf("./%s/oapi", config["gen-dir"]), 0755)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
