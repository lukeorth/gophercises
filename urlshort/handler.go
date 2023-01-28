package urlshort

import (
    "encoding/json"
    "io"
    "net/http"

    "gopkg.in/yaml.v2"
)

type shortUrl struct {
    Path    string 
    Url     string
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if url, ok := pathsToUrls[r.URL.Path]; ok {
            http.Redirect(w, r, url, http.StatusFound)
        }
        fallback.ServeHTTP(w, r)
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
func YAMLHandler(yml io.Reader, fallback http.Handler) (http.HandlerFunc, error) {
    parsedYaml, err := parseYAML(yml)
    if err != nil {
        return fallback.ServeHTTP, nil
    }
    pathMap := buildMap(parsedYaml)
    return MapHandler(pathMap, fallback), nil
}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements the http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
//
// JSON is expected to be in the format:
//
//      [
//          {
//              "path": "/some-path",
//              "url": "https://www.some-url.com/demo
//          }
//      ]
//
// The only errors that can be returned all related to having
// invalid JSON data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func JSONHandler(jsn io.Reader, fallback http.Handler) (http.HandlerFunc, error) {
    parsedJson, err := parseJSON(jsn)
    if err != nil {
        return fallback.ServeHTTP, nil
    }
    pathMap := buildMap(parsedJson)
    return MapHandler(pathMap, fallback), nil
}

func buildMap(urls []shortUrl) map[string]string {
    m := make(map[string]string)
    for _, v := range urls {
        m[v.Path] = v.Url
    }
    return m
}

func parseYAML(yml io.Reader) ([]shortUrl, error) {
    s := []shortUrl{}
    err := yaml.NewDecoder(yml).Decode(&s)
    if err != nil {
        return nil, err
    }
    return s, nil
}

func parseJSON(jsn io.Reader) ([]shortUrl, error) {
    s := []shortUrl{}
    err := json.NewDecoder(jsn).Decode(&s)
    if err != nil {
        return nil, err
    }
    return s, nil
}
