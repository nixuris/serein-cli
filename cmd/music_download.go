package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var musicDownloadCmd = &cobra.Command{
	Use:   "download [youtube-url]",
	Short: "Download audio from YouTube using yt-dlp",
	Long:  `Download audio from YouTube using yt-dlp with --extract-audio --embed-thumbnail --add-metadata -o "%(title)s.%(ext)s"`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		youtubeURL := args[0]

		ytDlpCmd := exec.Command("yt-dlp", "--extract-audio", "--embed-thumbnail", "--add-metadata", "-o", "%(title)s.%(ext)s", youtubeURL)
		ytDlpCmd.Stdout = os.Stdout
		ytDlpCmd.Stderr = os.Stderr

		if err := ytDlpCmd.Run(); err != nil {
			fmt.Println("Error downloading audio:", err)
			os.Exit(1)
		}
	},
}

func init() {
	musicCmd.AddCommand(musicDownloadCmd)
}
