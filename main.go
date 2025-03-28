package main

import (
	"github.com/clanflare/cf-prober-cli/cmd"
)

func main() {
	cmd.RootCommand.AddCommand(cmd.InitCommand)
	cmd.RootCommand.AddCommand(cmd.LoginCommand)

	cmd.Execute()
}
