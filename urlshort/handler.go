package main

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleMap(w, r, pathsToUrls, fallback)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
type yamlPathMap struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var yamlPaths []yamlPathMap
	yamlMap := make(map[string]string, len(yamlPaths))
	err := yaml.Unmarshal(yml, &yamlPaths)
	if err != nil {
		return nil, err
	}
	for _, path := range yamlPaths {
		yamlMap[path.Path] = path.Url
	}

	return func(w http.ResponseWriter, r *http.Request) {
		handleMap(w, r, yamlMap, fallback)
	}, nil
}

func handleMap(w http.ResponseWriter, r *http.Request, pathMap map[string]string, fallback http.Handler) {
	url := pathMap[r.URL.Path]
	if url != "" {
		http.Redirect(w, r, url, 301)
	}
	fallback.ServeHTTP(w, r)
}
