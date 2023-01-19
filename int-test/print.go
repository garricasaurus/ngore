package main

import (
	"encoding/json"
	"fmt"
)

func PrintJson(i interface{}) {
	bytes, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
