package commit

import (
	"fmt"
	"strings"
)

var validCommitTypes = []string{"feat", "fix", "chore", "docs", "test", "refactor", "style", "build", "perf", "ci"}

func FormatCommitMessage(commitType, scope, message string) string {
	if message == "" {
		return fmt.Sprintf("%s: %s", commitType, scope)
	}
	return fmt.Sprintf("%s(%s): %s", commitType, scope, message)
}

func IsValidCommitType(commitType string) bool {
	for _, v := range validCommitTypes {
		if v == commitType {
			return true
		}
	}
	return false
}

func ParseArgs(args []string) (string, string) {
	if strings.Contains(args[0], ":") {
		return "", strings.Join(args, " ")
	}
	return args[0], strings.Join(args[1:], " ")
}

func PrintValidCommitTypes() {
	fmt.Println("Valid commit types:", strings.Join(validCommitTypes, ", "))
}

func ExtractFlags(args []string) []string {
	flags := []string{}
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			flags = append(flags, arg)
		}
	}
	return flags
}
