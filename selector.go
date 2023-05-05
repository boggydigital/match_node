package match_node

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
)

const (
	idPrefix    = "#"
	classPrefix = "."
)

type selector struct {
	et atom.Atom
	id string
	cl string
}

func NewSelector(query string) Matcher {
	tag, id, class := "", "", ""

	if strings.Contains(query, idPrefix) {
		rest := ""
		tag, rest, _ = strings.Cut(query, idPrefix)
		if strings.Contains(rest, classPrefix) {
			id, class, _ = strings.Cut(rest, classPrefix)
		} else {
			id = rest
		}
	} else if strings.Contains(query, classPrefix) {
		tag, class, _ = strings.Cut(query, classPrefix)
	} else {
		tag = query
	}

	a := atom.Lookup([]byte(tag))

	return &selector{
		et: a,
		id: id,
		cl: class,
	}
}

func (s *selector) Match(node *html.Node) bool {

	if (s.et != 0 && node.DataAtom != s.et) ||
		(s.id != "" && len(node.Attr) == 0) ||
		(s.cl != "" && len(node.Attr) == 0) {
		return false
	}

	if s.id != "" {
		for _, attr := range node.Attr {
			if attr.Key == "id" && attr.Val == s.id {
				return true
			}
		}
	}

	if s.cl != "" {
		for _, attr := range node.Attr {
			if attr.Key == "class" && strings.Contains(attr.Val, s.cl) {
				return true
			}
		}
	}

	return false
}
