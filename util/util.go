package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Debug(obj interface{}) {
	data, _ := json.Marshal(obj)
	fmt.Println(string(data))
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func NormalizePath(path string) string {
	dir, _ := filepath.Abs(path)
	return dir
}
