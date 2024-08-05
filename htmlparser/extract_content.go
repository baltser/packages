package htmlparser

import "golang.org/x/net/html"

func (p *DefaultHTMLParser) ExtractContent() []string {
	var textContent []string
	var inRubric bool

	p.forEachNode(p.RootNode, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == *p.ParentNode {
			if value, ok := p.attr(n, *p.AttrKey); ok && value == *p.AttrVal {
				inRubric = true
			} else {
				inRubric = false
			}
		}
		if inRubric && n.Type == html.ElementNode && n.Data == *p.LastChild {
			text := p.text(n)
			textContent = append(textContent, text)
		}
	})
	return textContent
}
