package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type testCase struct {
	path        string
	status      int
	redirectUrl string
}

func TestYamlHandlerWithFallback(t *testing.T) {
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	testCases := []testCase{
		{
			path:        "/urlshort",
			status:      301,
			redirectUrl: "https://github.com/gophercises/urlshort",
		},
		{
			path:        "/google-it",
			status:      200,
			redirectUrl: "",
		},
		{
			path:        "/urlshort-godoc",
			status:      301,
			redirectUrl: "https://godoc.org/github.com/gophercises/urlshort",
		},
	}
	mapHandler := MapHandler(pathsToUrls, defaultMux())
	handler, err := YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		t.Fatal(err)
	}
	runTest(handler, testCases, t)

}
func TestYamlHandler(t *testing.T) {
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	handler, err := YAMLHandler([]byte(yaml), defaultMux())
	if err != nil {
		t.Fatal(err)
	}
	testCases := []testCase{
		{
			path:        "/urlshort",
			status:      301,
			redirectUrl: "https://github.com/gophercises/urlshort",
		},
		{
			path:        "/urlshort-godoc",
			status:      200,
			redirectUrl: "",
		},
		{
			path:        "/google-it",
			status:      200,
			redirectUrl: "",
		},
	}
	runTest(handler, testCases, t)
}

func TestMapHandler(t *testing.T) {
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	testCases := []testCase{
		{
			path:        "/urlshort-godoc",
			status:      301,
			redirectUrl: "https://godoc.org/github.com/gophercises/urlshort",
		},
		{
			path:        "/google-it",
			status:      200,
			redirectUrl: "",
		},
	}
	handler := MapHandler(pathsToUrls, defaultMux())
	runTest(handler, testCases, t)
}

func runTest(handler http.Handler, testCases []testCase, t *testing.T) {
	for _, c := range testCases {
		func(tc testCase) {
			rr := httptest.NewRecorder()
			req, err := http.NewRequest("GET", tc.path, nil)
			if err != nil {
				t.Fatal(err)
			}
			handler.ServeHTTP(rr, req)

			// Check the status code is what we expect.
			if rr.Result().StatusCode != tc.status {
				t.Errorf("Response code does not match, expected: %d, actual: %d", tc.status, rr.Result().StatusCode)
			}
			if tc.redirectUrl != "" {
				responseUrl, err := rr.Result().Location()
				if err != nil {
					t.Fatal("error getting location")
				}
				responseLocation := responseUrl.String()
				if responseLocation != tc.redirectUrl {
					t.Errorf("Response redirect does not match. expected %s, actual: %s", tc.redirectUrl, responseLocation)
				}
			}
		}(c)
	}

}
