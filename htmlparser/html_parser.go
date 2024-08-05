package htmlparser

import (
	"bytes"
	"golang.org/x/net/html"
)

func (p *DefaultHTMLParser) attr(n *html.Node, key string) (string, bool) {
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val, true
		}
	}
	return "", false
}

func (p *DefaultHTMLParser) text(n *html.Node) string {
	var buf bytes.Buffer
	p.forEachNode(n, func(n *html.Node) {
		if n.Type == html.TextNode {
			buf.WriteString(n.Data)
		}
	})
	return buf.String()
}

func (p *DefaultHTMLParser) forEachNode(nodes *html.Node, f func(*html.Node)) {
	var traverse func(n *html.Node)
	traverse = func(n *html.Node) {
		f(n)
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}
	traverse(nodes)
}
