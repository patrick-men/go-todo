package functions

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

type Todo struct {
	Task     string `json:"task"`
	Priority string `json:"priority"`
	DueDate	 string	`json:"dueDate"`
}

func ListCheck() (b bool) {

	// Get user 
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting user's home directory:", err)
		return
	}

	dirPath := filepath.Join(usr.HomeDir, ".config", "todo")
	filePath := dirPath + "/default.json"
	stateFile := dirPath + "/.state.json"

	defaultTodoListJson := []Todo{
		{"Fix something", "M", ""},
		{"Change this to that", "H", "23.10.2023"},
		{"Answer mail", "L", "Today"},
	}

	// Check if the default todo list file exists
	_, err = os.Stat(filePath)
	if err != nil {

		
		// Create directory for tool
		err = os.MkdirAll(dirPath, 0755)
		if err != nil {
			fmt.Println("Error getting user's home directory:", err)
			return
		}

		// Create first default list
		state, err := os.Create(stateFile)
		if err != nil {
			fmt.Println("Error creating state file:", err)
			return
		}

		// Create state file, and write "default" to it
		data := []byte("default.json\n")
		err = os.WriteFile(stateFile, data, 0644)
		if err != nil {
			fmt.Println("Error writing to state file:", err)
			return
		}
		state.Close()

		// Create first default list
		file, err := os.Create(filePath)
		fmt.Println("Since this is your first time using ToDo, we have created a default list.")
		if err != nil {
			log.Fatal(err)
		}

		// Encode default values
		jsonData, err := json.MarshalIndent(defaultTodoListJson, "", "    ")
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}

		// Write to file
		err = os.WriteFile(filePath, jsonData, 0644)
		if err != nil {
			fmt.Println("Error writing JSON to file:", err)
			return
		}

		// Close the file
		file.Close()
		b = true
	} else {
		b = false
	}
	return b
}