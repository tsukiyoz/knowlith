package app

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	defaultHomeDir    = ".knowlith"
	defaultConfigName = "apiserver.yaml"
)

func searchDirs() []string {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return []string{filepath.Join(homeDir, defaultHomeDir), "."}
}

func filePath() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return filepath.Join(home, defaultHomeDir, defaultConfigName)
}
