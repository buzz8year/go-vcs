package command

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	addcmd "go-vcs/command/add"
	config "go-vcs/command/config"
	logcmd "go-vcs/command/log"
	"go-vcs/file"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
)

func Run(args []string) {
	if len(args) > 1 {
		Create(strings.Join(args[1:], " "))
	} else {
		fmt.Println("Message was not passed.")
	}
}

func Create(message string) {
	trackedFiles := addcmd.GetTrackedFiles()
	lastCommitFiles := GetLastCommitFiles()

	if len(trackedFiles) < 1 {
		fmt.Println("Nothing to commit.")
		return
	}

	if len(lastCommitFiles) > 0 {
		isModified := false
		for i, f := range lastCommitFiles {
			lastCommitHash, _ := file.CalculateHash(f)
			isModified = file.IsModified(addcmd.GetTrackedFiles()[i], lastCommitHash)
		}
		if !isModified {
			fmt.Println("Nothing to commit.")
			return
		}
	}

	commitId := GenerateId()
	CreateCommitFolder(commitId)
	CopyCommitFiles(commitId)

	logcmd.Append(commitId, message)
	fmt.Println("Changes are committed.")
}

func GenerateId() string {
	data := fmt.Sprintf("%d", time.Now().UnixNano())
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)

	hash := sha256.Sum256(append([]byte(data), randomBytes...))
	fullString := hex.EncodeToString(hash[:])

	return fullString[:40]
}

func CreateCommitFolder(commitId string) {
	commitDir := filepath.Join(config.CommitDir, commitId)
	if err := os.MkdirAll(commitDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func CopyCommitFiles(commitId string) {
	for _, f := range addcmd.GetTrackedFiles() {
		err := file.CopyTo(f, filepath.Join(config.CommitDir, commitId))
		if err != nil {
			return
		}
	}
}

func GetCommitEntries(commitId string) []os.DirEntry {
	commitDir := filepath.Join(config.CommitDir, commitId)
	entries, err := os.ReadDir(commitDir)
	if err != nil {
		log.Fatal(err)
	}
	return entries
}

func GetCommitFileByBase(commitId, basename string) (string, bool) {
	commitDir := filepath.Join(config.CommitDir, commitId)

	for _, entry := range GetCommitEntries(commitId) {
		if entry.Name() == basename {
			return filepath.Join(commitDir, entry.Name()), true
		}
	}
	return "", false
}

func GetCommitFileBases(commitId string) []string {
	var bases []string
	for _, entry := range GetCommitEntries(commitId) {
		bases = append(bases, entry.Name())
	}
	slices.Sort(bases)
	return bases
}

func GetCommitFiles(commitId string) []string {
	var files []string
	commitDir := filepath.Join(config.CommitDir, commitId)

	for _, entry := range GetCommitEntries(commitId) {
		files = append(files, filepath.Join(commitDir, entry.Name()))
	}
	slices.Sort(files)
	return files
}

func GetLastCommitFiles() []string {
	lastLog := logcmd.GetLast()
	if lastLog == "" {
		return []string{}
	}

	split := strings.Split(lastLog, " ")
	return GetCommitFiles(split[1])
}
