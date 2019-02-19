// Application to generate NodeJS code for CRUD operations
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	in := []byte(`{ 
		"name": "todo",
		"fields" : [
			{"name": "id", "type" : "int"},
			{"name": "title", "type" : "string"}
		]
	}`)
	var raw map[string]interface{}
	json.Unmarshal(in, &raw)
	fmt.Println(raw["name"].(string))

	fields := raw["fields"].([]interface{})

	for i := 0; i < len(fields); i++ {
		field := fields[i].(map[string]interface{})
		fmt.Println(field["name"].(string), "=>", field["type"].(string))
	}

	// Checking the command line arguments
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	if os.Args[1] == "-v" || os.Args[1] == "--version" {
		showVersion()
		return
	}

	showUsage()
}

func showVersion() {
	fmt.Println("Version 1.0")
}

func showUsage() {
	fmt.Println("Usage: node-cg <filename>")
}
