package music

import (
	"github.com/spf13/cobra"
	"serein/internal/shared"
)

var YTMusicDownloadCmd = shared.NewCommand(
	"download [youtube-url]",
	"Download audio from YouTube using yt-dlp",
	cobra.ExactArgs(1),
	func(cmd *cobra.Command, args []string) {
		youtubeURL := args[0]
		shared.ExecuteCommand(
			"yt-dlp",
			"--extract-audio",
			"--embed-thumbnail",
			"--add-metadata",
			"-o", "%(title)s.%(ext)s",
			youtubeURL,
		)
	},
)
