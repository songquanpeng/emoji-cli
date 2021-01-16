package main

import (
	"fmt"
	"log"
)

func printHelpInfo()  {
	fmt.Println("Help information:")
	fmt.Println(`1. emoji description: search by description.
2. emoji -h: print help information.
3. emoji -u: update local database.`)
}

func updateData()  {
	fetchData()
	updateDatabase()
}

func fetchData()  {
	log.Print("Fetching data from Github.")
	log.Println("Data fetched.")
}

func updateDatabase()  {
	log.Print("Updating local database.")

	log.Println("Database updated.")
}

func processQuery(description string)  {
	fmt.Println(description)
}

func searchDatabase(description string)  {

}

func printItem()  {

}