package add

import (
	"bufio"
	"fmt"
	"go-vcs/command"
	config "go-vcs/command/config"
	"go-vcs/file"
	"log"
	"os"
	"path/filepath"
	"slices"
)

func Run(args []string) {
	if len(args) > 1 {
		AppendTrackedFile(args[1])
	} else {
		PrintTrackedFiles()
	}
}

func AppendTrackedFile(filePath string) {
	if !file.Exists(filePath) {
		fmt.Printf("Can't find '%s'.\n", filePath)
		return
	}

	files := GetTrackedFiles()
	for _, f := range files {
		if f == filePath {
			fmt.Printf("The file '%s' is tracked.\n", filePath)
			return
		}
	}

	file.AppendLine(filepath.Join(config.TargetDir, "index.txt"), filePath)
	fmt.Printf("The file '%s' is tracked.\n", filePath)
}

func GetTrackedFiles() []string {
	var files []string

	f, err := os.Open(filepath.Join(config.TargetDir, "index.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		files = append(files, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	slices.Sort(files)
	return files
}

func PrintTrackedFiles() {
	files := GetTrackedFiles()

	if len(files) < 1 {
		fmt.Println(command.Get("add").Description)
	} else {
		fmt.Println("Tracked files:")
		for _, f := range files {
			fmt.Println(f)
		}
	}
}
