package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Debug(obj ...interface{}) {
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
