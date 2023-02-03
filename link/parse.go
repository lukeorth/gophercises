package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)


type Link struct {
    Href string
    Text string
}

func Parse(r io.Reader) ([]Link, error) {
    var links []Link

    doc, err := html.Parse(r)
    if err != nil {
        return nil, err
    }
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
            href := n.Attr[0].Val
            text := strings.TrimSpace(n.FirstChild.Data)

            links = append(links, Link{Href: href, Text: text})
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
    f(doc)

    return links, err
}
