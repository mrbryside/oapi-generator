package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "oapi-generator",
	Short: "A code generator for OpenAPI specifications",
	Long:  `oapi-generator is a command-line tool that generates Go code from an OpenAPI specification. Use it to create robust, type-safe API clients and servers with minimal effort.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to oapi-generator!")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
