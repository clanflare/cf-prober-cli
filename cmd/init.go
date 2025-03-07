package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/clanflare/cf-pariksha-cli/types"
	"github.com/clanflare/cf-pariksha-cli/utils"
	"github.com/cqroot/prompt"
	"github.com/spf13/cobra"
)

var InitCommand = &cobra.Command{
	Use:     "initialize",
	Aliases: []string{"init"},
	Short:   "Initialize a cli config",
	Long:    `Initialize a cli config`,
	Run: func(cmd *cobra.Command, args []string) {

		cwd, cwdErr := os.Getwd()

		if cwdErr != nil {
			log.Println("Error while getting current working directory")
			log.Println(cwdErr)
			return
		}

		if utils.CheckFileExists(cwd + "./manifest.json") {
			log.Println("Manifest file already exists in the current directory")
		}

		content := types.Manifest{
			Token:      "",
			Client_Uri: "",
			Repos:      map[string]string{},
			Projects:   []types.Project{},
		}

		clientUriValue, _ := prompt.New().Ask("Enter the client uri").Input("")

		content.Client_Uri = clientUriValue

		jsonData, jsonErr := json.MarshalIndent(content, "", "  ")

		if jsonErr != nil {
			log.Println("Error while marshalling json")
			log.Println(jsonErr)
			return
		}

		utils.CreateFile(cwd+"/manifest.json", jsonData)
	},
}
