package fasturlchecker

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

// FastChecker will check simply faster than slowchecker
func FastChecker() {
	results := make(map[string]string)
	c := make(chan requestResult)
	urls := []string{
		"https://www.google.com/",
		"https://flatuicolors.com/",
		"https://github.com/",
		"https://www.facebook.com/",
		"https://www.amazon.com/",
		"https://mycolor.space/",
		"https://instagram.com/",
		"https://airbnb.com",
		"https://soundcloud.com",
	}
	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan requestResult) {
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		c <- requestResult{url: url, status: "FAILED"}
	} else {
		c <- requestResult{url: url, status: "OK"}
	}
}
