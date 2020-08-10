package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func main() {
	args := os.Args[1:]
	fileName := args[0]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Errorf("error opening file")
	}
	links := parseHtml(html.NewTokenizer(file))
	for _, link := range *links {
		fmt.Printf("%+v\n", link)
	}
}
func parseHtml(tokenizer *html.Tokenizer) *[]Link {
	links := make([]Link, 0)
	for {
		var href *string
		var text *string
		token := tokenizer.Next()
		if token == html.ErrorToken {
			return &links
		}
		if token == html.StartTagToken {
			tagName, _ := tokenizer.TagName()
			if string(tagName) == "a" {
			Attr:
				for {
					attr, val, more := tokenizer.TagAttr()
					attrName := string(attr)
					if attrName == "href" {
						h := string(val)
						href = &h
						break Attr
					}
					if !more {
						break Attr
					}
				}
				var t []byte
			Text:
				for {
					next := tokenizer.Next()
					t = append(t, tokenizer.Text()...)
					if next == html.EndTagToken {
						tagName, _ := tokenizer.TagName()
						if string(tagName) == "a" {
							t = append(t, tokenizer.Text()...)
							break Text
						}
					}
				}
				tString := strings.TrimSpace(string(t))
				text = &tString
				if href != nil && text != nil {
					links = append(links, Link{Text: *text, Href: *href})
				}
			}
		}
	}
}
