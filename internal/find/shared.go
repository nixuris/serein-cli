package find

import (
	"fmt"
	"os"

	"serein/internal/shared"
)

func DeletePath(path string, isDir bool) {
	var err error
	if isDir {
		err = os.RemoveAll(path)
	} else {
		err = os.Remove(path)
	}
	if err != nil {
		fmt.Printf("❌ Failed to delete %s: %v\n", path, err)
	} else {
		fmt.Printf("✅ Deleted: %s\n", path)
	}
}

func DeleteGrepMatches(basePath, term string) {
	matches, err := shared.RunCommand("grep", "-rlE", term, basePath)
	if err != nil {
		fmt.Printf("Error finding files to delete: %v\n", err)
		return
	}
	for _, match := range matches {
		DeletePath(match, false)
	}
}
