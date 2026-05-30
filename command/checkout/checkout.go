package command

import (
	"fmt"
	addcmd "go-vcs/command/add"
	commit "go-vcs/command/commit"
	config "go-vcs/command/config"
	"go-vcs/file"
	"log"
	"os"
	"path/filepath"
	"slices"
)

func Run(args []string) {
	if len(args) < 2 {
		fmt.Println("Commit id was not passed.")
	} else {
		commitDir := filepath.Join(config.TargetDir, "commits", args[1])

		if file.Exists(commitDir) {
			SwitchToCommit(args[1])
			fmt.Printf("Switched to commit %s.\n", args[1])
		} else {
			fmt.Println("Commit does not exist.")
		}
	}
}

func SwitchToCommit(commitId string) {
	commitFileBases := commit.GetCommitFileBases(commitId)
	trackedFiles := addcmd.GetTrackedFiles()

	for _, f := range trackedFiles {

		basename := filepath.Base(f)

		if !slices.Contains(commitFileBases, basename) {
			_ = os.Remove(f)
			continue
		}

		commitFile, exists := commit.GetCommitFileByBase(commitId, basename)
		if !exists {
			continue
		}

		data, err := os.ReadFile(commitFile)
		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(f, data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}
