// Application to generate NodeJS code for CRUD operations
package main

import (
	"fmt"
	"node-cg/model"
	"os"
)

func main() {

	inputObject, _ := model.BuildInput(`{
		"name": "todo",
		"fields" : [
			{"name": "id", "type" : "int"},
			{"name": "title", "type" : "string"}
		],
		"settings" : {
			"list" : "true",
			"add" : true,
			"details" : true,
			"update" : true,
			"delete" : true
		}
	}`)

	fmt.Println(inputObject)

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
