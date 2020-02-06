package slowurlchecker

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	errRequestFail = errors.New("Requeset Failed")
)

// LoopUrls mollaa
func SlowChecker() {
	var results = make(map[string]string)
	// if you don't use make() or var smth = map[string]string{}
	// the variable is gonna be "nil". Not an empty map.
	urls := []string{
		"https://www.google.com/",
		"https://flatuicolors.com/",
		"https://github.com/",
		"https://www.facebook.com/",
		"https://www.amazon.com/",
		"https://mycolor.space/",
		"https://instagram.com/",
		"https://airbnb.com",
	}
	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
			fmt.Println(err)
		} else {
			fmt.Println("Checking : " + url + " " + result)
		}
		results[url] = result
	}
	fmt.Println(results)
}

func hitURL(url string) error {
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		return errRequestFail
	}
	return nil
}
