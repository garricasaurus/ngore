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
		sr.Health = extractHealth(node)
		sr.Peers = extractPeers(node)
		sr.Seeds = extractSeeds(node)
		sr.Size = extractSize(node)
		sr.Uploaded = extractUploaded(node)
		sr.Uploader = extractUploader(node)
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

func extractTitle(n *html.Node) string {
	a := parse.GetFirstChildWithTagName(n, "a")
	if a != nil {
		return titleAttr(a)
	}
	return ""
}

func extractAltTitle(n *html.Node) string {
	nodes := parse.GetElementsByClass(n, "siterank")
	if len(nodes) == 0 {
		return ""
	}
	span := parse.GetFirstChildWithTagName(nodes[0], "span")
	if span == nil {
		return ""
	}
	return titleAttr(span)
}

func extractHealth(n *html.Node) string {
	nodes := parse.GetElementsByClass(n, "box_d2")
	if len(nodes) == 0 {
		return ""
	}
	return parse.GetText(nodes[0])
}

func extractPeers(n *html.Node) string {
	nodes := parse.GetElementsByClass(n, "box_l2")
	if len(nodes) == 0 {
		return ""
	}
	a := parse.GetFirstChildWithTagName(nodes[0], "a")
	if a == nil {
		return ""
	}
	return parse.GetText(a)
}

func extractSeeds(n *html.Node) string {
	nodes := parse.GetElementsByClass(n, "box_s2")
	if len(nodes) == 0 {
		return ""
	}
	a := parse.GetFirstChildWithTagName(nodes[0], "a")
	if a == nil {
		return ""
	}
	return parse.GetText(a)
}

func extractSize(n *html.Node) string {
	nodes := parse.GetElementsByClass(n, "box_meret2")
	if len(nodes) == 0 {
		return ""
	}
	return parse.GetText(nodes[0])
}

func extractUploaded(n *html.Node) string {
	nodes := parse.GetElementsByClass(n, "box_feltoltve2")
	if len(nodes) == 0 {
		return ""
	}
	return parse.GetText(nodes[0])
}

func extractUploader(n *html.Node) string {
	nodes := parse.GetElementsByClass(n, "box_feltolto2")
	if len(nodes) == 0 {
		return ""
	}
	spans := parse.GetElementsByClass(nodes[0], "feltolto_szin")
	if len(spans) == 0 {
		return ""
	}
	return parse.GetText(spans[0])
}

func titleAttr(element *html.Node) string {
	title, ok := parse.FindAttr(element, "title")
	if ok {
		return title
	}
	return ""
}
