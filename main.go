package main

import (
	"log"
	"os"
)

func main() {
	accessToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoiYWFrYXNoIiwiZXhwIjoxNzI4ODMwMjYxfQ.R7GGwa9Poe2ug7nLJy7ULDkTFWcNf05CoIhvvPnJQfs"
	if len(os.Args) == 2 {
		if len(os.Args[1]) == len(accessToken) {
			accessToken = os.Args[1]
		} else {
			log.Fatal("inconsitent access token")
		}
	}
	htmlContent := extractData(accessToken)

	printExtracted(htmlContent)
}
