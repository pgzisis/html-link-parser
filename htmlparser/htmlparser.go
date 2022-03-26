package htmlparser

import (
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func newLink(href, text string) Link {
	return Link{Href: href, Text: text}
}

// Parse returns a list of links out of an html file
func Parse(file io.Reader) []Link {
	doc, err := html.Parse(file)
	if err != nil {
		log.Fatalln(err)
	}

	var links []Link

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					href, text := strings.TrimSpace(a.Val), strings.TrimSpace(n.FirstChild.Data)
					link := newLink(href, text)
					links = append(links, link)

					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	return links
}
