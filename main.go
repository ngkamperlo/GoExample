package main

import (
	"net/http"

	"github.com/naimis20/GoExample/scraper"
	"github.com/naimis20/GoExample/server"
)

func main() {
	// pages, err := strconv.Atoi(os.Args[1])
	// if err != nil {
	// 	fmt.Println("Error")
	// 	pages = 3
	// }

	pages := 3

	client := &http.Client{}

	server.Serve(scraper.ScrapeBooking(client, pages))
}
