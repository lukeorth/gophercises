package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/lukeorth/gophercises/link"
)

/*
   1. GET the webpage
   2. parse all the links on the page
   3. build proper urls with our links
   4. filter our any links w/ a diff domain
   5. find all pages (BFS)
   6. print out XML
*/

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
    Value string `xml:"loc"`
}

type urlset struct {
    Urls []loc `xml:"url"`
    Xmlns string `xml:"xmlns,attr"`
}

func main() {
    urlFlag := flag.String("url", "https://www.calhoun.io", "the url that you want to build a sitemap for")
    maxDepth := flag.Int("depth", 1, "the maximum number of links deep to traverse")
    flag.Parse()

    pages := bfs(*urlFlag, *maxDepth)
    toXml := urlset{
        Xmlns: xmlns,
    }
    for _, p := range pages {
        toXml.Urls = append(toXml.Urls, loc{p})
    }
    fmt.Print(xml.Header)
    enc := xml.NewEncoder(os.Stdout)
    enc.Indent("", "  ")
    if err := enc.Encode(toXml); err != nil {
        panic(err)
    }
    fmt.Println()
}

func bfs(urlStr string, maxDepth int) []string {
    seen := make(map[string]bool)
    var q map[string]bool
    nq := map[string]bool{
        urlStr: true,
    }
    for i := 0; i <= maxDepth; i++ {
        q, nq = nq, make(map[string]bool)
        for url := range q {
            if _, ok := seen[url]; ok {
                continue
            }
            seen[url] = true
            for _, l := range get(url) {
                nq[l] = true 
            }
        }
    }
    ret := make([]string, 0, len(seen))
    for url := range seen {
        ret = append(ret, url)
    }
    return ret
}

func get(urlStr string) []string {
    resp, err := http.Get(urlStr)
    if err != nil {
        return []string{}
    }
    defer resp.Body.Close()

    reqUrl := resp.Request.URL
    baseUrl := &url.URL {
        Scheme: reqUrl.Scheme,
        Host:   reqUrl.Host,
    }
    base := baseUrl.String()
    return filter(hrefs(resp.Body, base), withPrefix(base))
}

func hrefs(r io.Reader, base string) []string {
    links, _ := link.Parse(r)
    var ret []string
    for _, l := range links {
        switch {
        case strings.HasPrefix(l.Href, "/"):
            ret = append(ret, base + l.Href)
        case  strings.HasPrefix(l.Href, "http"):
            ret = append(ret, l.Href)
        }
    }
    return ret
}

func filter(links []string, keepFn func(string) bool) []string {
    var ret []string
    for _, l := range links {
        if keepFn(l) {
            ret = append(ret, l)
        }
    }
    return ret
}

func withPrefix(pfx string) func(string) bool {
    return func(link string) bool {
        return strings.HasPrefix(link, pfx)
    }
}
