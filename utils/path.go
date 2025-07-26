package utils

import (
	"log"
	"os"
	"path/filepath"
)

func GetAppDataPath(extraPaths ...string) string {
	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		dataHome = filepath.Join(homeDir, ".local", "share")
	}
	finalPaths := append([]string{dataHome, "dottest"}, extraPaths...)
	return filepath.Join(finalPaths...)
}
