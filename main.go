package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func download_page(url string) (io.Reader, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

type HTMLMeta struct {
	Title    string
	SiteName string `json:"site_name"`
}

// Helper to get the meta informations of a page
// From https://gist.github.com/inotnako/c4a82f6723f6ccea5d83c5d3689373dd
func extract_meta(content io.Reader) *HTMLMeta {
	tokenizer := html.NewTokenizer(content)

	titleFound := false

	hm := new(HTMLMeta)

	for {
		tt := tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			return hm
		case html.StartTagToken, html.SelfClosingTagToken:
			t := tokenizer.Token()
			if t.Data == `body` {
				return hm
			}
			if t.Data == "title" {
				titleFound = true
			}
			if t.Data == "meta" {
				ogTitle, ok := extractMetaProperty(t, "og:title")
				if ok {
					hm.Title = ogTitle
				}

				ogSiteName, ok := extractMetaProperty(t, "og:site_name")
				if ok {
					hm.SiteName = ogSiteName
				}
			}
		case html.TextToken:
			if titleFound {
				t := tokenizer.Token()
				hm.Title = t.Data
				titleFound = false
			}
		}
	}
	return hm
}

func extractMetaProperty(token html.Token, property string) (content string, ok bool) {
	for _, attr := range token.Attr {
		if attr.Key == "property" && attr.Val == property {
			ok = true
		}
		if attr.Key == "content" {
			content = attr.Val
		}
	}
	return
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
	attributes := extract_meta(content)
	fmt.Println(attributes)
}
