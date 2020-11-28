package main

import (
	"fmt"
	"net/http"
)

func main() {
	content, err := contentType("https://golang.org")
	if err != nil {
		fmt.Printf("ERROR %s \n", err)
	} else {
		fmt.Printf("content is %s \n", content)
	}
}

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	content := resp.Header.Get("Content-Type")
	if content == "" {
		return "", fmt.Errorf("Unable to find content type header")
	}

	return content, nil
}

func close() {

}
