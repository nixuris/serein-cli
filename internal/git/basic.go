package git

import (
	"github.com/spf13/cobra"
)

func RegisterBasicGitCommands(parent *cobra.Command) {
    parent.AddCommand(gitSyncCmd)
    parent.AddCommand(gitStageCmd)
    parent.AddCommand(gitUnstageCmd)
    parent.AddCommand(gitUndoCmd)
    parent.AddCommand(gitChangesCmd)
    parent.AddCommand(gitStatusCmd)
}

var gitSyncCmd = &cobra.Command{
	Use:   "sync [branch]",
	Short: "Alias for git pull origin <branch>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("pull", "origin", args[0])
	},
}

var gitStageCmd = &cobra.Command{
	Use:   "stage [path]",
	Short: "Alias for git add <path>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("add", args[0])
	},
}

var gitUnstageCmd = &cobra.Command{
	Use:   "unstage [path]",
	Short: "Alias for git reset <path>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("reset", args[0])
	},
}

var gitUndoCmd = &cobra.Command{
	Use:   "undo [path]",
	Short: "Alias for git restore <path>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("restore", args[0])
	},
}

var gitChangesCmd = &cobra.Command{
	Use:   "changes [path]",
	Short: "Alias for git diff <path>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("diff", args[0])
	},
}

var gitStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Alias for git status",
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("status")
	},
}
