package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"example.com/GoExample/scraper"
	"example.com/GoExample/server"
)

func main() {
	pages, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error")
		pages = 3
	}

	client := &http.Client{}

	server.Serve(scraper.ScrapeBooking(client, pages))
}