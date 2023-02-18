package cmd

import (
	"fmt"
	"github.com/mrbryside/oapi-generator/generator"
	"github.com/mrbryside/oapi-generator/reader"
	"github.com/mrbryside/oapi-generator/validator"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate-server",
	Short: "generate server via oapi-codegen",
	Long:  `generate server by oapi-codegen`,
	RunE: func(cmd *cobra.Command, args []string) error {
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		config := reader.ReadConfig()
		result := validator.Config(name, config)
		if result != "" {
			return fmt.Errorf(result)
		}

		// create generated folder
		err = generator.GenerateFolder(config)
		if err != nil {
			return fmt.Errorf("error running init folder: %v", err)
		}

		// generate server by oapi-codegen
		err = generator.GenerateServer(name, *config["gen-dir"], *config["spec-dir"])
		if err != nil {
			fmt.Printf("error generating server files: %v", err)
		}
		fmt.Println("Generate server successfully!!")
		return nil
	}}

func init() {
	generateCmd.Flags().String("name", "", "Name of spec file (name.yaml)")
	rootCmd.AddCommand(generateCmd)
}
