package util

import (
	"os"
	"path/filepath"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func NormalizePath(path string) string {
	dir, _ := filepath.Abs(path)
	return dir
}
