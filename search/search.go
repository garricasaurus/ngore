package search

import (
	"git.okki.hu/garric/ngore/parse"
	"golang.org/x/net/html"
)

func ParseResponse(doc *html.Node) *Result {
	return &Result{
		Torrents: parseTorrents(doc),
		Page:     parsePageInfo(doc),
	}
}

func parseTorrents(doc *html.Node) []*Torrent {
	torrents := make([]*Torrent, 0)
	nodes := parse.GetElementsByClass(doc, "box_torrent")
	for _, node := range nodes {
		t := &Torrent{}
		txt := getTxtNode(node)
		if txt != nil {
			t.Title = extractTitle(txt)
			t.AltTitle = extractAltTitle(txt)
		}
		t.Health = extractHealth(node)
		t.Peers = extractPeers(node)
		t.Seeds = extractSeeds(node)
		t.Size = extractSize(node)
		t.Uploaded = extractUploaded(node)
		t.Uploader = extractUploader(node)
		torrents = append(torrents, t)
	}
	return torrents
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
