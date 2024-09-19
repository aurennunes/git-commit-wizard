package main

import (
	"fmt"
	"os"

	"github.com/aurennunes/git-commit-wizard/internal/commit"
	"github.com/aurennunes/git-commit-wizard/internal/git"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: git commit-wizard <type> [<scope>] <subject> [-a | -amend]")
		commit.PrintValidCommitTypes()
		return
	}

	commitType := os.Args[1]
	scope, message := commit.ParseArgs(os.Args[2:])

	commitMessage := commit.FormatCommitMessage(commitType, scope, message)
	flags := commit.ExtractFlags(os.Args)

	git.RunCommit(commitMessage, flags)
}
