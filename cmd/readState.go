package cmd

import (
    "fmt"
    "os"
    "os/user"
    "path/filepath"
)

var currentState string

func readState() error {
	usr, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to get current user: %v", err)
	}

	stateFilePath := filepath.Join(usr.HomeDir, ".config/todo/.state.json")
	content, err := os.ReadFile(stateFilePath)
	if err != nil {
		return fmt.Errorf("error reading state file: %v", err)
	}

	currentState = filepath.Base(string(content))
	return nil
}

func getState() {
    // Initialize the state when Cobra commands are executed
    if err := readState(); err != nil {
        fmt.Println("error reading state file:", err)
        os.Exit(0)
    }
}