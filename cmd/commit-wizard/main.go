package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: git commit-wizard <type> [<scope>] <subject> [-a | -amend]")

		return
	}
}
