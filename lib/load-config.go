package lib

import (
	"encoding/json"
	"os"

	"github.com/clanflare/cf-prober-cli/types"
	"github.com/clanflare/cf-prober-cli/utils"
)

func LoadConfig() types.Manifest {
	cwd, cwdErr := os.Getwd()

	if cwdErr != nil {
		panic("Error getting current working directory")
	}

	if !utils.CheckFileExists(cwd + "/manifest.json") {
		utils.Red.Println("Manifest not found!")
		utils.Yellow.Println("Please run `cf-prober-cli init` to create a new project.")
		os.Exit(0)
	}

	configBytes, configError := utils.ReadFileBytes(cwd + "/manifest.json")

	if configError != nil {
		utils.Red.Println("Error reading config file")
		os.Exit(0)
	}

	manifest := types.Manifest{}

	jsonError := json.Unmarshal(configBytes, &manifest)

	if jsonError != nil {
		utils.Red.Println("Error un-marshalling JSON data")
		os.Exit(0)
	}

	return manifest
}
