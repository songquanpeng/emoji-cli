package main

import (
	"os"
	"path"
	"strings"
)
var dataDir = ".emoji-cli"
var jsonPath = "./data.json"
var databasePath = "./data.db"

func initialize()  {
	home, _ := os.UserHomeDir()
	dataDir = path.Join(home, dataDir)
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		_ = os.Mkdir(dataDir, 0777)
	}
	jsonPath = path.Join(dataDir, jsonPath)
	if _, err := os.Stat(jsonPath); os.IsNotExist(err) {
		fetchData()
	}
	databasePath = path.Join(dataDir, databasePath)
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
