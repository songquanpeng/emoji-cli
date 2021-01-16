package main

import (
	"os"
	"strings"
)

var dataDir = ".emoji-cli"
var jsonPath = "./data.json"
var databasePath = "./data.db"

func initialize() {
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		_ = os.Mkdir(dataDir, 0777)
	}
	if _, err := os.Stat(jsonPath); os.IsNotExist(err) {
		fetchData()
	}
	var count int64
	db.Table("emojis").Count(&count)
	if count == 0 {
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
	} else {
		executeCommand(strings.Join(os.Args[1:], " "))
	}
}
