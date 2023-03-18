package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func GenerateFolder(config map[string]*string) error {
	// create folder in generated path
	err := os.MkdirAll(fmt.Sprintf("./%s/oapi", *config["gen-dir"]), 0755)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func GenerateServer(name, genPath, specPath string) error {
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
	dtoCmd := exec.Command("oapi-codegen", "-generate", "types", "-o", fmt.Sprintf("%s/oapi/%sdto/%sdto.go", genPath, name, name), "-package", fmt.Sprintf("%sdto", name), "--generate-types-pointer=false", fmt.Sprintf("%s/%s.yaml", specPath, name))
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
