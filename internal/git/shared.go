package git

import (
	"serein/internal/shared"
)

func runGitCommand(args ...string) {
	shared.ExecuteCommand("git", args...)
}
