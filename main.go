// Application to generate NodeJS code for CRUD operations
package main

import (
	"fmt"
	"os"
)

func main() {

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
