package main

import (
	"encoding/json"
	"fmt"
	"github.com/flatten_unflatten/converting"
)

func main() {
	var m interface{}
	defaultStr := `{
		"1": "3",
			"5": {
			"8": "jhsdgflksdmf",
				"67": {
				"4": "ksjdghf"
			}
		}
	}`
	json.Unmarshal([]byte(defaultStr), &m)

	flatten, err := converting.Flatten(m.(map[string]interface{}))
	if err != nil {
		fmt.Println("cannot create flatten map")
	}
	fmt.Println("Flatten map: ", flatten)

	unflatten, err := converting.Unflatten(m.(map[string]interface{}))
	if err != nil {
		fmt.Println("cannot create unflatten map")
	}
	fmt.Println("Unflatten map: ", unflatten)
}
