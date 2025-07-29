package utils

import (
	"path/filepath"
)

func GetAppDataPath(extraPaths ...string) string {
	dataHome := "/etc/"
	finalPaths := append([]string{dataHome, "dottest"}, extraPaths...)
	return filepath.Join(finalPaths...)
}
