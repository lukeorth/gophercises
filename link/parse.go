package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="...">) in an HTML
// document.
type Link struct {
    Href string
    Text string
}

// Parse will take in an HTML document and will return a
// slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
    doc, err := html.Parse(r)
    if err != nil {
        return nil, err
    }
    /*
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
            href := n.Attr[0].Val
            text := strings.TrimSpace(n.FirstChild.Data)

            links = append(links, Link{href, text})
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
    f(doc)
    */
    nodes := linkNodes(doc)
    var links []Link
    for _, node := range nodes {
        links = append(links, buildLink(node))
    }
    return links, err
}

func buildLink(n *html.Node) Link {
    var ret Link
    for _, attr := range n.Attr {
        if attr.Key == "href" {
            ret.Href = attr.Val
            break
        }
    }
    ret.Text = text(n)
    return ret
}

func text(n *html.Node) string {
    if n.Type == html.TextNode {
        return n.Data
    }
    if n.Type != html.ElementNode {
        return ""
    }
    var ret string
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        ret += strings.TrimSpace(text(c) + " ")
    }
    return ret
}

func linkNodes(n *html.Node) []*html.Node {
    if n.Type == html.ElementNode && n.Data == "a" {
        return []*html.Node{n}
    }
    var ret []*html.Node
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        ret = append(ret, linkNodes(c)...)
    }
    return ret
}
