package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

func download_page(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func main() {
	url := flag.String("url", "", "the url of the page to get the informations of")
	flag.Parse()

	if *url == "" {
		log.Fatal("url is required")
	}
	content, err := download_page(*url)
	if err != nil {
		log.Fatalf("error while trying to get the content of page %s: %s\n", *url, err)
	}
	fmt.Println(content)
}
