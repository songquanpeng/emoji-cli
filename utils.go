package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func printHelpInfo() {
	fmt.Println("Help information:")
	fmt.Println(`1. emoji description: search by description.
2. emoji -h: print help information.
3. emoji -u: update local database.

Notice: database is saved at ` + databasePath)
}

func updateData() {
	fetchData()
	updateDatabase()
}

func fetchData() {
	log.Print("Fetching data from Github.")
	url := "https://raw.githubusercontent.com/songquanpeng/emoji-cli/main/data.json"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	out, err := os.Create(jsonPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	log.Println("Data fetched.")
}

func updateDatabase() {
	log.Print("Updating local database.")
	UpdateDatabase()
	log.Println("Database updated.")
}

func processQuery(description string) {
	keywords := strings.Split(description, " ")
	emojis := Query(keywords[0])
	for i, emoji := range emojis {
		fmt.Printf("%d. %s :%s:\n", i+1, emoji.Emoji, emoji.Shortcode)
	}
}
