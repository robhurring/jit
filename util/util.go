package util

import (
	"encoding/json"
	"fmt"
)

func Debug(obj interface{}) {
	data, _ := json.Marshal(obj)
	fmt.Println(string(data))
}
