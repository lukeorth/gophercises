package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lukeorth/gophercises/cyoa"
)

func main() {
    port := flag.Int("port", 3000, "the port on which to start the CYOA web application")
    filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
    flag.Parse()
    fmt.Printf("Using the story in %s.\n", *filename)

    f, err := os.Open(*filename)
    if err != nil {
        panic(err)
    }

    story, err := cyoa.JsonStory(f)
    if err != nil {
        panic(err)
    }

    h := cyoa.NewHandler(story)
    fmt.Printf("Starting the server at: %d\n", *port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
