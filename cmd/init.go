package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var InitCommand = &cobra.Command{
	Use:     "initialize",
	Aliases: []string{"init"},
	Short:   "Initialize a cli config",
	Long:    `Initialize a cli config`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Initializing config...")
	},
}
