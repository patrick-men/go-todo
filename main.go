// main.go
package main

import (
	"os"
	"todo/cmd"
	"todo/functions"
)

func main() {

	// If default file is created, exit script
	if functions.ListCheck(){
		os.Exit(0)
	}
	
	cmd.Execute()
}
