package search

import (
	"git.okki.hu/garric/ncgore/parse"
	"golang.org/x/net/html"
)

func ParseResults(doc *html.Node) ([]*Result, error) {
	result := make([]*Result, 0)
	nodes := parse.GetElementsByClass(doc, "box_torrent")
	for _, node := range nodes {
		sr := &Result{}
		txt := getTxtNode(node)
		if txt != nil {
			sr.Title = extractTitle(txt)
			sr.AltTitle = extractAltTitle(txt)
		}
		result = append(result, sr)
	}
	return result, nil
}

func getTxtNode(n *html.Node) *html.Node {
	txtNodes := parse.GetElementsByClass(n, "torrent_txt")
	if len(txtNodes) == 0 {
		return nil
	}
	return txtNodes[0]
}

func extractTitle(txt *html.Node) string {
	a := parse.GetFirstChildWithTagName(txt, "a")
	if a != nil {
		return titleAttr(a)
	}
	return ""
}

func extractAltTitle(txt *html.Node) string {
	srNodes := parse.GetElementsByClass(txt, "siterank")
	if len(srNodes) == 0 {
		return ""
	}
	span := parse.GetFirstChildWithTagName(srNodes[0], "span")
	if span == nil {
		return ""
	}
	return titleAttr(span)
}

func titleAttr(element *html.Node) string {
	title, ok := parse.FindAttr(element, "title")
	if ok {
		return title
	}
	return ""
}
