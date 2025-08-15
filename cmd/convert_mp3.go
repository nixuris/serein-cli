package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var convertMp3NewCmd = &cobra.Command{
	Use:   "mp3 [directory]",
	Short: "Convert opus/flac to mp3",
	Long:  `Convert opus/flac to mp3 using ffmpeg.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		directory := args[0]

		logFile, err := os.Create(filepath.Join(directory, "conversion_errors.log"))
		if err != nil {
			fmt.Println("Error creating log file:", err)
			os.Exit(1)
		}
		defer logFile.Close()

		err = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && (strings.HasSuffix(info.Name(), ".flac") || strings.HasSuffix(info.Name(), ".opus")) {
				base := strings.TrimSuffix(path, filepath.Ext(path))
				outfile := base + ".mp3"

				if _, err := os.Stat(outfile); err == nil {
					fmt.Println("Skipping (exists):", outfile)
					return nil
				}

				fmt.Println("Converting + embedding cover:", path, "->", outfile)
				ffmpegCmd := exec.Command("ffmpeg", "-nostdin", "-i", path, "-map", "0:a", "-map", "0:v?", "-c:a", "libmp3lame", "-q:a", "0", "-id3v2_version", "3", "-metadata:s:v", "title=Album cover", "-metadata:s:v", "comment=Cover (front)", outfile)

				var stderr strings.Builder
				ffmpegCmd.Stderr = &stderr

				if err := ffmpegCmd.Run(); err != nil {
					logFile.WriteString(fmt.Sprintf("Conversion error for %s: %v\nFFmpeg Output:\n%s\n", path, err, stderr.String()))
					fmt.Printf("Conversion error for %s: %v\nFFmpeg Output:\n%s\n", path, err, stderr.String())
					return nil // Continue walking even if one conversion fails
				}

				// Verify output file size
				if info, err := os.Stat(outfile); err == nil && info.Size() > 0 {
					fmt.Println("Converted:", outfile)
					fmt.Println("Deleting source:", path)
					os.Remove(path)
				} else {
					logFile.WriteString(fmt.Sprintf("Conversion failed (zero-size or missing output): %s\n", path))
					fmt.Printf("Conversion failed (zero-size or missing output): %s\n", path)
				}
			}
			return nil
		})

		if err != nil {
			fmt.Println("Error during file walk:", err)
			os.Exit(1)
		}

		fmt.Println("All done! Check conversion_errors.log for any errors.")
	},
}

func init() {
	convertCmd.AddCommand(convertMp3NewCmd)
}
