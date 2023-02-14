package match_node

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type elementTypeId struct {
	et atom.Atom
	id string
}

func NewId(elementType atom.Atom, id string) Matcher {
	return &elementTypeId{
		et: elementType,
		id: id,
	}
}

func (eti *elementTypeId) Match(node *html.Node) bool {
	if node.DataAtom != eti.et ||
		(eti.id != "" && len(node.Attr) == 0) {
		return false
	}

	for _, attr := range node.Attr {
		if attr.Key == "id" {
			return attr.Val == eti.id
		}
	}

	return false
}
