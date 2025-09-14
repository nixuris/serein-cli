package git

import (
	"github.com/spf13/cobra"
	"serein/internal/shared"
)

func BasicGitCommands(parent *cobra.Command) {
	parent.AddCommand(gitSyncCmd)
	parent.AddCommand(gitAddRemoteCmd)
	parent.AddCommand(gitStageCmd)
	parent.AddCommand(gitUnstageCmd)
	parent.AddCommand(gitUndoCmd)
	parent.AddCommand(gitChangesCmd)
	parent.AddCommand(gitStatusCmd)
}

var gitSyncCmd = shared.NewCommand(
	"sync [branch]",
	"Alias for git pull origin <branch>",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		runGitCommand("pull", "origin", args[0])
	},
)

var gitAddRemoteCmd = shared.NewCommand(
	"remote [url]",
	"Alias for git remote add origin <url>",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		runGitCommand("remote", "add", "origin", args[0])
	},
)

var gitStageCmd = shared.NewCommand(
	"stage [paths...]",
	"Alias for git add <paths>",
	cobra.MinimumNArgs(1),
	func(cmd *cobra.Command, args []string) {
		gitArgs := append([]string{"add"}, args...)
		runGitCommand(gitArgs...)
	},
)

var gitUnstageCmd = shared.NewCommand(
	"unstage [path]",
	"Alias for git reset <path>",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		runGitCommand("reset", args[0])
	},
)

var gitUndoCmd = shared.NewCommand(
	"undo [path]",
	"Alias for git restore <path>",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		runGitCommand("restore", args[0])
	},
)

var gitChangesCmd = shared.NewCommand(
	"changes [path]",
	"Alias for git diff <path>",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		runGitCommand("diff", args[0])
	},
)

var gitStatusCmd = shared.NewCommand(
	"status",
	"Alias for git status",
	cobra.ExactArgs(0),
	func(cmd *cobra.Command, args []string) {
		runGitCommand("status")
	},
)
