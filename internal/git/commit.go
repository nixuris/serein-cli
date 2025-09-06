package git

import (
	"github.com/spf13/cobra"
)

func init() {
	CommitCmd.AddCommand(CommitPushCmd)
	CommitCmd.AddCommand(CommitListCmd)
	CommitCmd.AddCommand(CommitUndoCmd)
	CommitCmd.AddCommand(CommitDeleteCmd)
	CommitCmd.AddCommand(CommitChangesCmd)
	CommitCmd.AddCommand(CommitCompareCmd)
}

var CommitCmd = &cobra.Command{
	Use:   "commit [message]",
	Short: "Alias for git commit -m \"<commit msg>\"",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("commit", "-m", args[0])
	},
}

var CommitPushCmd = &cobra.Command{
	Use:   "push [branch]",
	Short: "Alias for git push origin <branch>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("push", "origin", args[0])
	},
}

var CommitListCmd = &cobra.Command{
	Use:   "list",
	Short: "Alias for git log --oneline",
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("log", "--oneline")
	},
}

var CommitUndoCmd = &cobra.Command{
	Use:   "undo [SHA]",
	Short: "Alias for git revert <SHA>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("revert", args[0])
	},
}

var CommitDeleteCmd = &cobra.Command{
	Use:   "delete [number]",
	Short: "Alias for git reset --hard HEAD~<numb>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("reset", "--hard", "HEAD~"+args[0])
	},
}

var CommitChangesCmd = &cobra.Command{
	Use:   "changes [SHA]",
	Short: "Alias for git show <SHA>",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("show", args[0])
	},
}

var CommitCompareCmd = &cobra.Command{
	Use:   "compare [SHA1] [SHA2]",
	Short: "Alias for git diff <SHA1> <SHA2>",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand("diff", args[0], args[1])
	},
}
