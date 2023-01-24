package search

import (
	"git.okki.hu/garric/ngore/parse"
	"golang.org/x/net/html"
	"strconv"
	"strings"
)

const defaultPage = 1
const pageSize = 25

func parsePageInfo(n *html.Node) (pi *PageInfo) {
	pi = &PageInfo{}
	pager := parse.GetElementById(n, "pager_bottom")
	if pager != nil {
		pi.Current = parseCurrent(n)
		pi.Prev = parsePrev(n)
		pi.Next = parseNext(n)
	}
	return
}

func parseCurrent(n *html.Node) int {
	span := parse.GetElementByClass(n, "active_link")
	if span == nil {
		return 1
	}
	str := parse.GetElementByTag(span, "strong")
	return calcPageNumber(parse.GetText(str))
}

func parsePrev(n *html.Node) int {
	prev := parse.GetElementById(n, "pPa")
	if prev == nil {
		return parseCurrent(n)
	}
	str := parse.GetElementByTag(prev, "strong")
	return calcPageNumber(parse.GetText(str))
}

func parseNext(n *html.Node) int {
	next := parse.GetElementById(n, "nPa")
	if next == nil {
		return parseCurrent(n)
	}
	str := parse.GetElementByTag(next, "strong")
	return calcPageNumber(parse.GetText(str))
}

func calcPageNumber(s string) int {
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		return defaultPage
	}
	n, err := strconv.Atoi(parts[0])
	if err != nil {
		return defaultPage
	}
	return n/pageSize + 1
}
