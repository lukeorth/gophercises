package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/lukeorth/gophercises/link"
)

func main() {
    // gracefully handle errors and exit
    var err error
    defer func() {
        if err != nil {
            log.Fatalln(err)
        }
    }()

    // get command args
    url := flag.String("url", "", "url of the website to map")
    flag.Parse()

    // make request to website
    res, err := http.Get(*url)
    if err != nil {
        return
    }
    if res.StatusCode > 299 {
        log.Printf("Response failed with status code: %d\n", res.StatusCode)
        return
    }

    // parse links
    links, err := link.Parse(res.Body)
    if err != nil {
        return
    }
    fmt.Printf("%+v\n", links)
}
