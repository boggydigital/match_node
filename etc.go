package match_node

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
)

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
