package activity

import (
	"git.okki.hu/garric/ngore/parse"
	"golang.org/x/net/html"
)

func ParseResponse(doc *html.Node) *Info {
	info := &Info{}
	parseTable(doc, info)
	parseHistory(doc, info)
	return info
}

func parseHistory(doc *html.Node, info *Info) {
	torrents := parse.GetElementsByClass(doc, "hnr_torrents")
	for _, torrent := range torrents {
		_ = TorrentActivity{
			Name:  parseName(torrent),
			Start: parseStart(torrent),
		}
	}
}

func parseStart(node *html.Node) string {
	return ""
}

func parseName(node *html.Node) string {
	a := parse.GetElementByTag(node, "a")
	if a == nil {
		return ""
	}
	name, _ := parse.FindAttr(a, "title")
	return name
}

func parseTable(doc *html.Node, info *Info) {
	element := parse.GetElementByClass(doc, "fobox_tartalom")
	if element == nil {
		return
	}
	data := parse.GetElementsByClass(element, "dd")
	if len(data) < 9 {
		return
	}
	info.Rank = Rank{
		Daily:     parse.GetText(data[0]),
		Weekly:    parse.GetText(data[1]),
		Monthly:   parse.GetText(data[2]),
		PrevMonth: parse.GetText(data[3]),
	}
	info.CanDownload = parse.GetText(data[4])
	info.Stats = Stats{
		Current:     parse.GetText(data[5]),
		Allowed:     parse.GetText(data[6]),
		PenMonths:   parse.GetText(data[7]),
		PenTorrents: parse.GetText(data[8]),
	}
}
