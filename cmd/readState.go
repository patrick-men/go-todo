package cmd

import (
	"io/ioutil"
    "fmt"
    "os"
    "path/filepath"
)

var currentState string

func readState() error {
	content, err := ioutil.ReadFile("~/.config/todo/.state.json")
    if err != nil {
        return err
    }

    currentState = filepath.Base(string(content))
    return nil
}

func getState() {
    // Initialize the state when Cobra commands are executed
    if err := readState(); err != nil {
        fmt.Println("Error reading state file:", err)
        os.Exit(1)
    }
}