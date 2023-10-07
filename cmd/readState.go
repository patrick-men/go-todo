package cmd

import (
    "fmt"
    "os"
    "path/filepath"
)

var currentState string

func readState() error {
	content, err := os.ReadFile("~/.config/todo/.state.json")
    if err != nil {
        return err
    }

	currentState = filepath.Base(string(content))
	return nil
}

func GetState() {
    // Initialize the state when Cobra commands are executed
    if err := readState(); err != nil {
        fmt.Println("error reading state file: owo", err)
        os.Exit(0)
    }
}