package match_node

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
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

type elementTypeClass struct {
	et atom.Atom
	cl string
	eq bool
}

func NewEtc(elementType atom.Atom, class string, classEquals bool) Matcher {
	return &elementTypeClass{
		et: elementType,
		cl: class,
		eq: classEquals,
	}
}

func (etc *elementTypeClass) Match(node *html.Node) bool {
	if node.DataAtom != etc.et ||
		(etc.cl != "" && len(node.Attr) == 0) {
		return false
	}

	if etc.cl == "" {
		return true
	}

	for _, attr := range node.Attr {
		if attr.Key == "class" {
			if etc.eq {
				return attr.Val == etc.cl
			} else {
				return strings.Contains(attr.Val, etc.cl)
			}
		}
	}

	return false
}
