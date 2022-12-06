package htmlparser

import (
	"golang.org/x/net/html"
	"io"
	"log"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatalf("Failed to parse html")
	}
	links := make([]Link, 0)
	createLinks(doc, &links)

	return links, nil
}

func createLinks(rootNode *html.Node, links *[]Link) {
	if rootNode == nil {
		return
	}

	if rootNode.Type == html.ElementNode && rootNode.Data == "a" {
		for _, a := range rootNode.Attr {
			if a.Key == "href" {
				text := concatSiblingAndChildTextNodes(rootNode.FirstChild)
				*links = append(*links, Link{
					Href: a.Val,
					Text: strings.ReplaceAll(text, "\n", ""),
				})
			}
		}
	}
	createLinks(rootNode.NextSibling, links)
	createLinks(rootNode.FirstChild, links)
}

func concatSiblingAndChildTextNodes(rootNode *html.Node) string {
	if rootNode == nil {
		return ""
	}
	if rootNode.Type == html.TextNode {
		return strings.TrimSpace(rootNode.Data) + concatSiblingAndChildTextNodes(rootNode.FirstChild) + concatSiblingAndChildTextNodes(rootNode.NextSibling)
	} else {
		return concatSiblingAndChildTextNodes(rootNode.FirstChild) + concatSiblingAndChildTextNodes(rootNode.NextSibling)
	}
}
