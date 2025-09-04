package archive

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func ExpandTargets(targets []string) []string {
    for i, target := range targets {
        info, err := os.Stat(target)
        if err == nil && info.IsDir() {
            targets[i] = filepath.Join(target, "/*")
        }
    }
    return targets
}

func BuildArchiveCommand(archiveName string, targets []string, password string) *exec.Cmd {
    fileExt := filepath.Ext(archiveName)
    cmdArgs := []string{"a"}

    if password != "" {
        cmdArgs = append(cmdArgs, "-p"+password)
    }

    if fileExt == ".zip" {
        cmdArgs = append(cmdArgs, "-tzip")
    }

    cmdArgs = append(cmdArgs, archiveName)
    cmdArgs = append(cmdArgs, targets...)

    return exec.Command("7z", cmdArgs...)
}

func BuildExtractCommand(target string, password string) *exec.Cmd {
    cmdArgs := []string{"x"}
    if password != "" {
        cmdArgs = append(cmdArgs, "-p"+password)
    }
    cmdArgs = append(cmdArgs, target)
    return exec.Command("7z", cmdArgs...)
}

func RunWithOutput(cmd *exec.Cmd, errorMessage string) {
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Println(errorMessage, err)
        os.Exit(1)
    }
}
