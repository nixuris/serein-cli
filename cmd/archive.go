package cmd

import (
    	"serein/internal/archive"
		"github.com/spf13/cobra"
)

func init() {
    rootCmd.AddCommand(archiveCmd)
	archiveCmd.AddCommand(archive.UnzipCmd)
	archiveCmd.AddCommand(archive.ZipCmd)
}

var archiveCmd = &cobra.Command{
    Use:   "archive",
    Short: "Archive commands",
    Long:  `Commands for creating and extracting archives.`,
}
