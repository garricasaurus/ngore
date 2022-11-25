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
		txt := parse.GetElementsByClass(node, "torrent_txt")
		if len(txt) < 1 {
			continue
		}
		a := parse.GetFirstChild(txt[0], "a")
		if a != nil {
			title, ok := parse.FindAttr(a, "title")
			if ok {
				sr.Title = title
			}
		}
		result = append(result, sr)
	}
	return result, nil
}
