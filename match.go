package match_node

import "golang.org/x/net/html"

type MatchDelegate func(node *html.Node) bool

func Match(doc *html.Node, matchDelegate MatchDelegate) *html.Node {
	var f func(*html.Node) *html.Node
	f = func(n *html.Node) *html.Node {
		if matchDelegate(n) {
			return n
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			val := f(c)
			if val != nil {
				return val
			}
		}
		return nil
	}
	return f(doc)
}
