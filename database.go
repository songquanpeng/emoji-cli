package main

import (
	"encoding/json"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var db *gorm.DB

type Emoji struct {
	Emoji     string `json:"emoji"`
	Shortcode string `json:"shortcode"`
	Keywords  string `json:"keywords"`
}

func init() {
	home, _ := os.UserHomeDir()
	dataDir = path.Join(home, dataDir)
	jsonPath = path.Join(dataDir, jsonPath)
	databasePath = path.Join(dataDir, databasePath)
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		_ = os.Mkdir(dataDir, 0777)
	}

	var err error
	db, err = gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		panic("failed to open data.db")
	}
	_ = db.AutoMigrate(&Emoji{})
}

func UpdateDatabase() {
	err := db.Migrator().DropTable(&Emoji{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&Emoji{})
	if err != nil {
		log.Fatal(err)
	}
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	var emojis []Emoji
	err = json.Unmarshal(byteValue, &emojis)
	if err != nil {
		log.Fatal(err)
	}
	for _, emoji := range emojis {
		db.Create(emoji)
	}
}

func Query(keyword string) (emojis []Emoji) {
	db.Where("keywords like '%" + keyword + "%'").
		Find(&emojis)
	return
}
