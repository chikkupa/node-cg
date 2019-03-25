// Application to generate NodeJS code for CRUD operations
package main

import (
	"fmt"
	"io/ioutil"
	"node-cg/model"
	"os"
)

func main() {

	// Checking the command line arguments
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	if os.Args[1] == "-v" || os.Args[1] == "--version" {
		showVersion()
		return
	}

	inputContent, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	inputObject, err := model.BuildInput(string(inputContent))

	if err != nil {
		fmt.Println("JSON Error: ", err.Error())
		return
	}

	err = inputObject.GenerateModel()

	if err != nil {
		fmt.Println("JSON Error: ", err.Error())
		return
	}

	fmt.Println(inputObject)

	showUsage()
}

func showVersion() {
	fmt.Println("Version 1.0")
}

func showUsage() {
	fmt.Println("Usage: node-cg <filename>")
}
