package htmlparser

import "golang.org/x/net/html"

type HTMLParser interface {
	attr(n *html.Node, key string) (string, bool)
	text(n *html.Node) string
	forEachNode(root *html.Node, f func(*html.Node))
	ExtractContent() []string
}

type DefaultHTMLParser struct {
	RootNode   *html.Node
	ParentNode *string
	AttrKey    *string
	AttrVal    *string
	LastChild  *string
}

func NewDefaultHTMLParser(htmlPage *html.Node, parentNode, attrKey, attrVal, lastChild string) HTMLParser {
	return &DefaultHTMLParser{
		RootNode:   htmlPage,
		ParentNode: &parentNode,
		AttrKey:    &attrKey,
		AttrVal:    &attrVal,
		LastChild:  &lastChild,
	}
}
