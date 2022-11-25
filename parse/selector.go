package parse

import (
	"golang.org/x/net/html"
)

type matcher func(n *html.Node) bool

func GetElementById(n *html.Node, id string) *html.Node {
	return traverse(n, func(n *html.Node) bool {
		return hasId(n, id)
	})
}

func GetElementsByClass(n *html.Node, className string) []*html.Node {
	return getElements(n, func(n *html.Node) bool {
		return hasClass(n, className)
	})
}

func GetFirstChildWithTagName(n *html.Node, tag string) *html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && c.Data == tag {
			return c
		}
	}
	return nil
}

func FindAttr(n *html.Node, key string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func getElements(n *html.Node, p matcher) []*html.Node {
	result := make([]*html.Node, 0)
	collect(n, p, &result)
	return result
}

func collect(n *html.Node, p matcher, result *[]*html.Node) {
	if p(n) {
		*result = append(*result, n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		collect(c, p, result)
	}
}

func traverse(n *html.Node, p matcher) *html.Node {
	if p(n) {
		return n
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		r := traverse(c, p)
		if r != nil {
			return r
		}
	}
	return nil
}

func hasClass(n *html.Node, className string) bool {
	if n.Type == html.ElementNode {
		class, ok := FindAttr(n, "class")
		if ok && class == className {
			return true
		}
	}
	return false
}

func hasId(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		i, ok := FindAttr(n, "id")
		if ok && i == id {
			return true
		}
	}
	return false
}
