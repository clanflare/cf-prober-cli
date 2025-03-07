package cmd

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/clanflare/cf-prober-cli/lib"
	"github.com/clanflare/cf-prober-cli/types"
	"github.com/clanflare/cf-prober-cli/utils"
	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/write"
	"github.com/spf13/cobra"
)

type responseBodyData struct {
	Token string `json:"token"`
}

var LoginCommand = &cobra.Command{
	Use:     "login",
	Aliases: []string{"login"},
	Short:   "Login in cli",
	Long:    `Login in cli`,
	Run: func(cmd *cobra.Command, args []string) {
		manifest := lib.LoadConfig()

		if manifest.Token != "" {
			utils.Red.Println("Already logged in")
			os.Exit(0)
		}

		utils.Green.Println("Please follow the given url to get a token: " + manifest.Client_Uri + "/auth/login/cli")

		magicToken, _ := prompt.New().Ask("Enter magic token: ").Write("", write.WithHelp(true))

		requestJson, requestJsonError := json.Marshal(map[string]string{
			"token": magicToken,
		})

		if requestJsonError != nil {
			utils.Red.Println("Error while marshalling json")
			os.Exit(0)
		}

		response, apiErr := http.Post(manifest.Client_Uri+"/api/auth/login/cli", "application/json", bytes.NewBuffer(requestJson))

		if apiErr != nil {
			utils.Red.Println("Error making JSON request")
			os.Exit(0)
		}

		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			utils.Red.Println("Login failed with status: ", response.Status)
			os.Exit(0)
		}

		responseData := types.APIResponse[responseBodyData]{}

		if err := json.NewDecoder(response.Body).Decode(&responseData); err != nil {
			utils.Red.Println("Error decoding JSON response")
			os.Exit(0)
		}

		if !responseData.Success {
			utils.Red.Println("Login failed: ", responseData.Error)
			os.Exit(0)
		}

		manifest.Token = responseData.Data.Token

		lib.SaveConfig(manifest)

		utils.Green.Println("Login successful")
	},
}
