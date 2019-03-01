package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	c := &http.Client{
		Timeout: 1 * time.Second,
	}

	res, err := c.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Println("Something went wrong")
	}
}
