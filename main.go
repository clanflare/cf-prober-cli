package main

import (
	"github.com/clanflare/cf-pariksha-cli/cmd"
)

func main() {
	cmd.RootCommand.AddCommand(cmd.InitCommand)

	cmd.Execute()
}
