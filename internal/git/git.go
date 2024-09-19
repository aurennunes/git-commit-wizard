package git

import (
	"fmt"
	"os"
	"os/exec"
)

func RunCommit(commitMessage string, flags []string) {
	args := []string{"commit", "-m", commitMessage}

	if contains(flags, "-a") {
		args = append(args, "-a")
	}

	if contains(flags, "-amend") {
		args = append(args, "--amend")
	}

	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running git command: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Commit successfully made!")
}

func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
