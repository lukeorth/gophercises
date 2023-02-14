package main

import (
	"flag"
	"fmt"
	"log"
	"os"

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
    fname := flag.String("file", "", "name of HTML file to parse")
    flag.Parse()

    // open file
    f, err := os.Open(*fname)
    if err != nil {
        return
    }
    defer f.Close()

    // business logic
    links, err := link.Parse(f)
    if err != nil {
        return
    }
    fmt.Printf("%+v\n", links)
}
