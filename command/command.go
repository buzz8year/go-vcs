package command

import (
	"fmt"
	"slices"
)

type Command struct {
	Name        string
	Description string
}

var Commands = []Command{
	{"config", "Get and set a username."},
	{"add", "Add a file to the index."},
	{"log", "Show commit logs."},
	{"commit", "Save changes."},
	{"checkout", "Restore a file."},
}

func Get(name string) Command {
	return Commands[GetIndex(name)]
}

func GetIndex(name string) int {
	return slices.IndexFunc(Commands, func(cmd Command) bool {
		return cmd.Name == name
	})
}

func PrintAll() {
	fmt.Println("These are SVCS command:")
	for _, cmd := range Commands {
		fmt.Printf("%-10s %s\n", cmd.Name, cmd.Description)
	}
}

func PrintWrong(cmd string) {
	fmt.Printf("'%s' is not a SVCS command.\n", cmd)
}
