package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func Colorize(color, text string) string {
	if runtime.GOOS == "windows" {
		return text
	}
	colors := map[string]string{
		"red":    "\033[31m",
		"green":  "\033[32m",
		"blue":   "\033[34m",
		"cyan":   "\033[36m",
		"yellow": "\033[33m",
		"reset":  "\033[0m",
	}
	return colors[color] + text + colors["reset"]
}

func GetSystemInfo() string {
	return fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH)
}
func IsAdmin() bool {
	if runtime.GOOS != "windows" {
		return os.Geteuid() == 0
	}
	cmd := exec.Command("net", "session")
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run() == nil
}
