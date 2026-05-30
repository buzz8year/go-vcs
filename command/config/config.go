package command

import (
	"bufio"
	"fmt"
	"go-vcs/file"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var TargetDir string
var CommitDir string

func Run(args []string) {
	if len(args) > 1 {
		SetUsername(args[1])
	}
	PrintUsername()
}

func CreateSystemFiles() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	TargetDir = filepath.Join(currentDir, "vcs")
	CommitDir = filepath.Join(TargetDir, "commits")

	if err := os.MkdirAll(TargetDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(CommitDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	file.Create(TargetDir, "config.txt")
	file.Create(TargetDir, "index.txt")
	file.Create(TargetDir, "log.txt")
}

func SetUsername(username string) {
	data := []byte("username=" + username)
	err := os.WriteFile(filepath.Join(TargetDir, "config.txt"), data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func GetUsername() string {
	f, err := os.Open(filepath.Join(TargetDir, "config.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		items := strings.Split(scanner.Text(), "=")
		if items[0] == "username" {
			return items[1]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return ""
}

func PrintUsername() {
	username := GetUsername()
	if username == "" {
		fmt.Println("Please, tell me who you are.")
	} else {
		fmt.Printf("The username is %s.\n", username)
	}
}
