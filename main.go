package main

import (
	"os"
	"strings"
)

var jsonPath = "./data.json"
var databasePath = "./data.db"

func initialize()  {
	if _, err := os.Stat(jsonPath); os.IsNotExist(err) {
		fetchData()
	}
	if _, err := os.Stat(databasePath); os.IsNotExist(err) {
		updateDatabase()
	}
}

func executeCommand(command string) {
	switch command {
	case "":
		fallthrough
	case "-h":
		printHelpInfo()
	case "-u":
		updateData()
	default:
		processQuery(command)
	}
	return
}

func main() {
	initialize()
	if len(os.Args) == 1 {
		printHelpInfo()
	}else {
		executeCommand(strings.Join(os.Args[1:], ""))
	}
}
