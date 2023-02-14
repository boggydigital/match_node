package match_node

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type elementAtom struct {
	at atom.Atom
}

func NewAtom(at atom.Atom) Matcher {
	return &elementAtom{
		at: at,
	}
}

func (ea *elementAtom) Match(node *html.Node) bool {
	if node.DataAtom == ea.at {
		return true
	}
	return false
}
