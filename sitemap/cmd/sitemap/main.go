package main

import (
	"flag"
	"log"

    "github.com/lukeorth/gophercises/sitemap"
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

    // parse links
    links := sitemap.Build(*url)
    _ = links
}
