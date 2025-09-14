package git

import (
	"github.com/spf13/cobra"
	"serein/internal/shared"
)

func init() {
	CommitCmd.AddCommand(CommitPushCmd)
	CommitCmd.AddCommand(CommitListCmd)
	CommitCmd.AddCommand(CommitUndoCmd)
	CommitCmd.AddCommand(CommitDeleteCmd)
	CommitCmd.AddCommand(CommitChangesCmd)
	CommitCmd.AddCommand(CommitCompareCmd)
}

var CommitCmd = shared.NewCommand(
	"commit [message]",
	"Alias for git commit -m \"<commit msg>\"",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		runGitCommand("commit", "-m", args[0])
	},
)

var CommitPushCmd = shared.NewCommand(
	"push [branch]",
	"Alias for git push origin <branch>",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		runGitCommand("push", "origin", args[0])
	},
)

var CommitListCmd = shared.NewCommand(
	"list",
	"Alias for git log --oneline",
	cobra.NoArgs,
	func(cmd *cobra.Command, args []string) {
		runGitCommand("log", "--oneline")
	},
)

var CommitUndoCmd = shared.NewCommand(
	"undo [SHA]",
	"Alias for git revert <SHA>",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		runGitCommand("revert", args[0])
	},
)

var CommitDeleteCmd = shared.NewCommand(
	"delete [number]",
	"Alias for git reset --hard HEAD~<numb>",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		runGitCommand("reset", "--hard", "HEAD~"+args[0])
	},
)

var CommitChangesCmd = shared.NewCommand(
	"changes [SHA]",
	"Alias for git show <SHA>",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		runGitCommand("show", args[0])
	},
)

var CommitCompareCmd = shared.NewCommand(
	"compare [SHA1] [SHA2]",
	"Alias for git diff <SHA1> <SHA2>",
	cobra.ExactArgs(2),
	func(cmd *cobra.Command, args []string) {
		runGitCommand("diff", args[0], args[1])
	},
)

