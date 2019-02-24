package main

import (
	"log"
	"net/http"
	"time"
)

var (
	timeout = 1 * time.Second
)

func main() {
	c := &http.Client{
		Timeout: timeout,
	}

	res, err := c.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		log.Println("Something went wrong")
	}
}
