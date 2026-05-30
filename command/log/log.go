package command

import (
	"bufio"
	"fmt"
	config "go-vcs/command/config"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func Run(args []string) {
	PrintAll()
}

func Append(commitId string, message string) {
	logText := "commit " + commitId + " Author: " + config.GetUsername() + " " + message
	logFile, err := os.OpenFile(filepath.Join(config.TargetDir, "log.txt"), os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	_, err = fmt.Fprintln(logFile, logText)
	if err != nil {
		log.Fatal(err)
	}
}

func GetLast() string {
	logs := GetAll()
	if len(logs) > 0 {
		return GetAll()[0]
	}
	return ""
}

func GetAll() []string {
	var logs []string

	file, err := os.Open(filepath.Join(config.TargetDir, "log.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		logs = append(logs, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	slices.Reverse(logs)
	return logs
}

func PrintAll() {
	logs := GetAll()

	if len(logs) < 1 {
		fmt.Println("No commits yet.")
	} else {
		for _, l := range logs {
			s := strings.Split(l, " ")
			fmt.Printf("%s %s\n%s %s\n", s[0], s[1], s[2], s[3])
			fmt.Println(strings.Join(s[4:], " "))
			fmt.Println()
		}
	}
}
