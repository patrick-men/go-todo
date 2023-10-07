// main.go
package main

import (
	"todo/cmd"
	"todo/functions"
	"fmt"
)

func main() {

	fmt.Println("oops")

	// If default file is created, exit script
	functions.ListCheck()

	
	cmd.Execute()
}
