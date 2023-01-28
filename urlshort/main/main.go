package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/lukeorth/gophercises/urlshort"
)

func main() {
    json := flag.String("json", "", "a json file which maps urls to shortend paths")
    yaml := flag.String("yaml", "", "a yaml file which maps urls to shortend paths")
    flag.Parse()

    // Build the MapHandler using the mux as the fallback
    pathsToUrls := map[string]string{
	    "/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	    "/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
    }
    mapHandler := newMapHandler(pathsToUrls)
    yamlHandler := newYAMLHandler(*yaml, mapHandler)
    jsonHandler := newJSONHandler(*json, yamlHandler)

    fmt.Println("Starting the server on :8080")
    http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func newMapHandler(paths map[string]string) http.HandlerFunc {
    mux := defaultMux()
    return urlshort.MapHandler(paths, mux)
}

func newYAMLHandler(fname string, fallback http.HandlerFunc) http.HandlerFunc {
    if fname == "" {
        return fallback.ServeHTTP
    }
    f, err := os.Open(fname)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    yamlHandler, err := urlshort.YAMLHandler(f, fallback)
    if err != nil {
        panic(err)
    }

    return yamlHandler
}

func newJSONHandler(fname string, fallback http.HandlerFunc) http.HandlerFunc {
    if fname == "" {
        return fallback.ServeHTTP
    }
    f, err := os.Open(fname)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    jsonHandler, err := urlshort.JSONHandler(f, fallback)
    if err != nil {
        panic(err)
    }

    return jsonHandler
}
