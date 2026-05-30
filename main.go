package main

import (
	"go-vcs/command"
	addcmd "go-vcs/command/add"
	checkout "go-vcs/command/checkout"
	commit "go-vcs/command/commit"
	config "go-vcs/command/config"
	logcmd "go-vcs/command/log"
	"os"
)

func main() {
	args := os.Args[1:]
	config.CreateSystemFiles()

	if len(args) < 1 || args[0] == "--help" {
		command.PrintAll()
	} else if command.GetIndex(args[0]) < 0 {
		command.PrintWrong(args[0])
	} else {
		switch args[0] {
		case "config":
			config.Run(args)
		case "add":
			addcmd.Run(args)
		case "log":
			logcmd.Run(args)
		case "commit":
			commit.Run(args)
		case "checkout":
			checkout.Run(args)
		}
	}
}
