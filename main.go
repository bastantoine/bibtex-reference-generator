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
	Title                string
	Author               string
	ArticleAuthor        string
	OgUpdatedTime        string
	ArticlePublishedTime string
	ArticleModifiedTime  string
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
				author, ok := extractMetaProperty(t, "author", "name")
				if ok {
					hm.Author = author
				}

				articleAuthor, ok := extractMetaProperty(t, "article:author", "property")
				if ok {
					hm.ArticleAuthor = articleAuthor
				}

				ogTitle, ok := extractMetaProperty(t, "og:title", "property")
				if ok {
					hm.Title = ogTitle
				}

				articlePublishedTime, ok := extractMetaProperty(t, "article:published_time", "property")
				if ok {
					hm.ArticlePublishedTime = articlePublishedTime
				}

				articleModifiedTime, ok := extractMetaProperty(t, "article:modified_time", "property")
				if ok {
					hm.ArticleModifiedTime = articleModifiedTime
				}

				ogUpdatedTime, ok := extractMetaProperty(t, "og:updated_time", "property")
				if ok {
					hm.OgUpdatedTime = ogUpdatedTime
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

func extractMetaProperty(token html.Token, property, metaType string) (content string, ok bool) {
	for _, attr := range token.Attr {
		if (metaType == "property" && attr.Key == "property" && attr.Val == property) ||
			(metaType == "name" && attr.Key == "name" && attr.Val == property) {
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
	fmt.Printf("Title: %s\n", attributes.Title)
	fmt.Printf("Author: %s\n", attributes.Author)
	fmt.Printf("OgUpdatedTime: %s\n", attributes.OgUpdatedTime)
	fmt.Printf("ArticlePublishedTime: %s\n", attributes.ArticlePublishedTime)
	fmt.Printf("ArticleModifiedTime: %s\n", attributes.ArticleModifiedTime)
}
