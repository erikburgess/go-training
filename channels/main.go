package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//for i := 0; i < len(links); i++ {
	//	go checkLink(<-c, c)
	//}

	//for {
	//	go checkLink( <-c, c)
	//}

	for l := range c {
		// Blocking call on the main routine. No messages will be pulled from the channel
		//time.Sleep(10 * time.Second)
		//go checkLink(l, c)
		go func() {
			time.Sleep(10 * time.Second)
			checkLink(l, c)
		}()
	}

}

func checkLink(link string, c chan string) {
	//time.Sleep(10 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		c <- link
		return
	}
	fmt.Printf("Site up %v\n", link)
	c <- link
}
