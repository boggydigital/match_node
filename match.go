package match_node

import (
	"golang.org/x/net/html"
)

type Matcher interface {
	Match(node *html.Node) bool
}

func Match(doc *html.Node, delegate Matcher) *html.Node {
	matches := Matches(doc, delegate, 1)
	if len(matches) > 0 {
		return matches[0]
	}
	return nil
}

func Matches(doc *html.Node, delegate Matcher, limit int) []*html.Node {
	matches := make([]*html.Node, 0)

	var f func(*html.Node)
	f = func(n *html.Node) {

		if delegate.Match(n) {
			matches = append(matches, n)
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
			if len(matches) == limit {
				return
			}
		}
	}

	f(doc)

	return matches
}
