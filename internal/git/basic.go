package git

import (
	"github.com/spf13/cobra"
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

var gitSyncCmd = &cobra.Command{
	Use:   "sync [branch]",
	Short: "Alias for git pull origin <branch>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("pull", "origin", args[0])
	},
}

var gitAddRemoteCmd = &cobra.Command{
	Use:   "remote [url]",
	Short: "Alias for git remote add origin <url>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("remote", "add", "origin", args[0])
	},
}

var gitStageCmd = &cobra.Command{
	Use:   "stage [paths...]",
	Short: "Alias for git add <paths>",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gitArgs := append([]string{"add"}, args...)
		runGitCommand(gitArgs...)
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
