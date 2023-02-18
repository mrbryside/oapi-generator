package generator

import (
	"flag"
	"fmt"
	"os"
)

func Execute() {
	// Define a flag for the directory name
	var dirName string
	flag.StringVar(&dirName, "dir", "", "the name of the directory to create jaaa")

	// Parse the command-line arguments
	flag.Parse()

	// Check if the directory name was provided
	if dirName == "" {
		fmt.Println("Please provide a directory name using the -dir flag")
		os.Exit(1)
	}

	// Create the new directory
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print a message indicating success
	fmt.Printf("Created directory '%s'\n", dirName)
}
