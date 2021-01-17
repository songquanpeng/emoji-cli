package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func printHelpInfo() {
	fmt.Println("Help information:")
	fmt.Println(`1. emoji description: search by description.
2. emoji -h: print help information.
3. emoji -u: update local database.`)
}

func updateData() {
	fetchData()
	updateDatabase()
}

func fetchData() {
	fmt.Println("Fetching data from Github...")
	url := "https://raw.githubusercontent.com/songquanpeng/emoji-cli/main/data.json"
	backupUrl := "https://gitee.com/songquanpeng/emoji-cli/raw/main/data.json"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to fetch data from Github.")
		fmt.Println("Fetching data from Gitee...")
		resp, err = http.Get(backupUrl)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Failed to fetch data from Gitee.")
			fmt.Println("Abort.")
			return
		}
	}
	defer resp.Body.Close()
	out, err := os.Create(jsonPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	fmt.Println("Data fetched.")
}

func updateDatabase() {
	fmt.Println("Updating local database...")
	UpdateDatabase()
	fmt.Println("Done.")
}

func processQuery(description string) {
	keywords := strings.Split(description, " ")
	emojis := Query(keywords[0])
	for _, emoji := range emojis {
		fmt.Printf("%s :%s:\n", emoji.Emoji, emoji.Shortcode)
	}
}
