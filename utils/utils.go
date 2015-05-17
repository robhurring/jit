package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Dump(obj interface{}) {
	data, _ := json.Marshal(obj)
	fmt.Println(string(data))
}

func Debug(obj interface{}) {
	fmt.Printf("%#v", obj)
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func NormalizePath(path string) string {
	dir, _ := filepath.Abs(path)
	return dir
}

func ExpandPath(path string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir

	if path[:2] == "~/" {
		path = strings.Replace(path, "~", dir, 1)
	}

	return path
}

func WalkTree(path string, callback func(path string)) {
	fullPath := ExpandPath(path)
	files, _ := ioutil.ReadDir(fullPath)

	for _, dir := range files {
		if !dir.IsDir() {
			continue
		}

		callback(filepath.Join(fullPath, dir.Name()))
	}
}
