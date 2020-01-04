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
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	linkNodes := findLinkNodes(doc)

	var ret []Link
	for _, node := range linkNodes {
		ret = append(ret, buildLink(node))
	}

	return ret, nil
}

func findLinkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		if findHref(n) != "#" {
			return []*html.Node{n}
		}
	}

	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, findLinkNodes(c)...)
	}
	return ret
}

func buildLink(n *html.Node) Link {
	return Link{
		Href: findHref(n),
		Text: findText(n),
	}
}

func findHref(n *html.Node) string {
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			return attr.Val
		}
	}
	return ""
}

func findText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += findText(c)
	}

	return strings.Join(strings.Fields(text), " ")
}
