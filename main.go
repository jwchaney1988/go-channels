package main

// Import required packages
import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Create a slice of strings for the list of URLs
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	// Create a channel of type string
	c := make(chan string)
	// Initial check if links are currently up and working
	for _, link := range links {
		// Execute the checkLink function passing in the current link and channel
		go checkLink(link, c)
	}

	// Loop to check the list of URLs as returned through the channel
	for l := range c {
		// Never access the same var from a child routine so passing in link (l)
		go func(link string) {
			// Wait 5 seconds
			time.Sleep(5 * time.Second)
			// Execute the checkLink function passing in the current link and channel
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	// Sending out get request
	_, err := http.Get(link)
	// Checking if there was an error returned
	if err != nil {
		fmt.Println(link, "might be down")
		// Passing in link to the channel
		c <- link
		// Exiting out of the function
		return
	}

	fmt.Println(link, "is up")
	// Passing in link to the channel
	c <- link
}
