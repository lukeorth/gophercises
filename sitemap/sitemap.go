package sitemap

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/lukeorth/gophercises/link"
)

type Urls map[string]bool

func Build(url string) ([]link.Link) {
    urls := bfs(url) 
    for k := range urls {
        fmt.Printf("%v\n", k)
    }
    fmt.Println(len(urls))
    return getLinks(url) 
}

func bfs(url string) (Urls) {
    visited := Urls{url: true}
    domain := url
    links := getDomainLinks(url, domain)

    for i := 0; i < len(links); i++ {
        _, ok := visited[links[i]]
        if !ok {
            visited[links[i]] = true
            links = append(links, getDomainLinks(links[i], domain)...)
        }
    }
    return visited
}

func getLinks(url string) ([]link.Link) {
    res, err := http.Get(url)
    if err != nil {
        return nil
    }

    links, err := link.Parse(res.Body)
    if err != nil {
        return nil
    }
    return links
}

func getDomainLinks(url string, domain string) ([]string) {
    links := []string{}
    for _, l := range getLinks(url) {
        if strings.HasPrefix(l.Href, domain) {
            links = append(links, strings.TrimRight(l.Href, "/"))
        }
        if strings.HasPrefix(l.Href, "/") {
            links = append(links, domain + strings.TrimRight(l.Href, "/"))
        }
    }
    return links
}
