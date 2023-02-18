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
		var config map[string]*string
		err = yaml.Unmarshal(data, &config)
		if err != nil {
			log.Fatalf("failed to parse YAML data: %v", err)
		}
		if config["spec-dir"] == nil || config["gen-dir"] == nil {
			return fmt.Errorf("error oapi-cfg.yaml is wrong")
		}

		yamlFile = fmt.Sprintf("%s/server.cfg.yaml", *config["spec-dir"])
		if _, err := os.Stat(yamlFile); os.IsNotExist(err) {
			return fmt.Errorf("spec server config (%s) does not exist", yamlFile)
		}
		yamlFile = fmt.Sprintf("%s/%s.yaml", *config["spec-dir"], name)
		if _, err := os.Stat(yamlFile); os.IsNotExist(err) {
			return fmt.Errorf("spec file (%s) does not exist", yamlFile)
		}

		// Run the command "go run main.go init"
		err = createFolder()
		if err != nil {
			return fmt.Errorf("error running init folder: %v", err)
		}

		err = generateServer(name, *config["gen-dir"], *config["spec-dir"])
		if err != nil {
			fmt.Printf("error generating server files: %v", err)
		}
		fmt.Println("generated")
		return nil
	}}

func init() {
	generateCmd.Flags().String("name", "", "Name of the application")
	rootCmd.AddCommand(generateCmd)
}

func generateServer(name, genPath, specPath string) error {
	// Remove existing DTO and service folders, and create new ones
	err := os.RemoveAll(fmt.Sprintf("%s/oapi/%sdto", genPath, name))
	if err != nil {
		return fmt.Errorf("error removing DTO folder: %v", err)
	}
	err = os.RemoveAll(fmt.Sprintf("%s/oapi/%ssrv", genPath, name))
	if err != nil {
		return fmt.Errorf("error removing service folder: %v", err)
	}
	err = os.MkdirAll(fmt.Sprintf("%s/oapi/%sdto", genPath, name), os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating DTO folder: %v", err)
	}
	err = os.MkdirAll(fmt.Sprintf("%s/oapi/%ssrv", genPath, name), os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating service folder: %v", err)
	}

	// Create a temporary server config file and write to it
	serverCfgPath := fmt.Sprintf("%s/server-%s.cfg.yaml", specPath, name)
	serverCfgContents := fmt.Sprintf("#name%sdto\n", name)
	err = ioutil.WriteFile(serverCfgPath, []byte(serverCfgContents), 0644)
	if err != nil {
		return fmt.Errorf("error creating temporary server config file: %v", err)
	}

	// Generate the DTO files using oapi-gen
	dtoCmd := exec.Command("oapi-codegen", "-generate", "types", "-o", fmt.Sprintf("%s/oapi/%sdto/%sdto.go", genPath, name, name), "-package", fmt.Sprintf("%sdto", name), fmt.Sprintf("%s/%s.yaml", specPath, name))
	dtoCmd.Dir = "./"
	_, err = dtoCmd.CombinedOutput()
	if err != nil {
		os.Remove(serverCfgPath)
		return fmt.Errorf("error generating DTO files: %v", err)
	}

	// Generate the service files using oapi-gen and the temporary config file
	srvCmd := exec.Command("oapi-codegen", "--config", serverCfgPath, "-package", fmt.Sprintf("%ssrv", name), "-o", fmt.Sprintf("%s/oapi/%ssrv/%ssrv.go", genPath, name, name), fmt.Sprintf("%s/%s.yaml", specPath, name))
	_, err = srvCmd.CombinedOutput()
	if err != nil {
		os.Remove(serverCfgPath)
		return fmt.Errorf("error generating service files: %v", err)
	}

	// Clean up the temporary server config file
	err = os.Remove(serverCfgPath)
	if err != nil {
		return fmt.Errorf("error removing temporary server config file: %v", err)
	}

	return nil
}
