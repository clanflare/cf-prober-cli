package lib

import (
	"encoding/json"
	"os"

	"github.com/clanflare/cf-prober-cli/types"
	"github.com/clanflare/cf-prober-cli/utils"
)

func SaveConfig(manifest types.Manifest) {
	cwd, cwdErr := os.Getwd()

	if cwdErr != nil {
		panic("Error getting current working directory")
	}

	configBytes, configError := json.MarshalIndent(manifest, "", "  ")

	if configError != nil {
		utils.Red.Println("Error marshalling config file")
		os.Exit(0)
	}

	err := utils.WriteFile(cwd+"/manifest.json", configBytes)
	if err != nil {
		utils.Red.Println("Error writing config file")
		os.Exit(0)
	}
}
