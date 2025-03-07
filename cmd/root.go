package cmd

import (
	"log"
	"os"
	"slices"

	"github.com/clanflare/cf-pariksha-cli/utils"
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "bc",
	Short: "bc is a command line tool for interacting with all the projects of Clanflare",
	Long:  "bc is a command line tool for interacting with all the projects of Clanflare. It allows you to manage, run, configure and reports issues of a project from a single place. It does not matter where and how.",
	Run: func(cmd *cobra.Command, args []string) {
		if slices.Contains(args, "help") {
			cmd.Help()
			return
		}

		utils.Green.Println("Welcome to bc CLI")
		utils.Yellow.Println("Use `bc init` to kick off and make sure to do this in a fresh folder.")
	},
}

func Execute() {
	if err := RootCommand.Execute(); err != nil {
		log.Println("Error!!")
		os.Exit(1)
	}
}
