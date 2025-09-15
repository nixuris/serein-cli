package git

import (
	"serein/internal/shared"
)

func runGitCommand(args ...string) {
	shared.CheckErr(shared.ExecuteCommand("git", args...))
}
