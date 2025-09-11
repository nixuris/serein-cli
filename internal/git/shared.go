package git

import (
	"serein/internal/execute"
)

func runGitCommand(args ...string) {
	execute.ExecuteCommand("git", args...)
}
