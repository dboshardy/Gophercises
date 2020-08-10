package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"testing"
)

type testCase struct {
	Filename string
	Links    []Link
}

func TestLinkParsing(t *testing.T) {
	testCases := []testCase{
		{
			Filename: "ex2.html",
			Links: []Link{
				{Text: "Check me out on twitter", Href: "https://www.twitter.com/joncalhoun"},
				{Text: "Gophercises is on Github!", Href: "https://github.com/gophercises"},
			},
		},
		{
			Filename: "ex1.html",
			Links: []Link{
				{Text: "A link to another page", Href: "/other-page"},
			},
		},
	}
	for _, tc := range testCases {
		r, err := os.Open(tc.Filename)
		if err != nil {
			t.Fatalf("Error opening file %s", tc.Filename)
		}
		tokenizer := html.NewTokenizer(r)
		links := parseHtml(tokenizer)
		fmt.Printf("%+v\n", links)
		for i, link := range *links {
			expected := tc.Links[i]
			if link != expected {
				t.Errorf("expected: %+v\nactual: %+v\n\n", expected, link)
				t.Fail()
			}
		}
	}

}
