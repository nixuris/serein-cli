package music

import (
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var YTMusicDownloadCmd = &cobra.Command{
	Use:   "download [youtube-url]",
	Short: "Download audio from YouTube using yt-dlp",
	Long:  `Download audio from YouTube using yt-dlp with --extract-audio --embed-thumbnail --add-metadata -o "%(title)s.%(ext)s"`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		youtubeURL := args[0]
		ytDlp := exec.Command(
			"yt-dlp",
			"--extract-audio",
			"--embed-thumbnail",
			"--add-metadata",
			"-o", "%(title)s.%(ext)s",
			youtubeURL,
		)

		if err := RunCommand(ytDlp, "Error downloading audio:"); err != nil {
			os.Exit(1)
		}
	},
}
