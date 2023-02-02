package details

import (
	"git.okki.hu/garric/ngore/parse"
	"golang.org/x/net/html"
	"regexp"
	"strings"
)

var detailLinkRegex = regexp.MustCompile(`.*\?(http.*)`)

func ParseDetails(doc *html.Node) *Details {
	return &Details{
		Title:       parseTitle(doc),
		ReleaseYear: parseReleaseYear(doc),
		Director:    parseDirector(doc),
		Actors:      parseActors(doc),
		Country:     parseCountry(doc),
		Labels:      parseLabels(doc),
		ImdbRating:  parseImdbRating(doc),
		ImdbLink:    parseImdbLink(doc),
		Length:      parseLength(doc),
		OtherLink:   parseOtherLink(doc),
		CoverImage:  parseCoverImage(doc),
		OtherImages: parseOtherImages(doc),
	}
}

func parseTitle(n *html.Node) string {
	div := parse.GetElementByClass(n, "infobar_title")
	if div == nil {
		return ""
	}
	return strings.TrimSpace(parse.GetText(div))
}

func parseReleaseYear(n *html.Node) string {
	return parseTableData(n, "Megjelenés éve:")
}

func parseDirector(n *html.Node) string {
	return parseTableData(n, "Rendező:")
}

func parseActors(n *html.Node) string {
	return parseTableData(n, "Szereplők:")
}

func parseCountry(n *html.Node) string {
	return parseTableData(n, "Ország:")
}

func parseLabels(n *html.Node) string {
	el := parseTableElement(n, "Címkék:")
	if el == nil {
		return ""
	}
	refs := parse.GetElementsByTag(el, "a")
	sb := &strings.Builder{}
	for _, ref := range refs {
		label := parse.GetText(ref)
		if label != "" {
			sb.WriteString(label + " ")
		}
	}
	return strings.TrimSpace(sb.String())
}

func parseImdbRating(n *html.Node) string {
	return parseTableData(n, "IMDb értékelés:")
}

func parseImdbLink(n *html.Node) string {
	return parseLink(n, "IMDb link:")
}

func parseLength(n *html.Node) string {
	return parseTableData(n, "Hossz:")
}

func parseOtherLink(n *html.Node) string {
	return parseLink(n, "Egyéb link:")
}

func parseCoverImage(n *html.Node) string {
	div := parse.GetElementByClass(n, "torrent_leiras")
	if div == nil {
		return ""
	}
	img := parse.GetElementByTag(div, "img")
	if img == nil {
		return ""
	}
	src, ok := parse.FindAttr(img, "src")
	if !ok {
		return ""
	}
	return src
}

func parseOtherImages(n *html.Node) []string {
	res := make([]string, 0)
	div := parse.GetElementByClass(n, "fobox_tartalom")
	if div == nil {
		return res
	}
	el := parse.GetElementByTag(div, "center")
	if el == nil {
		return res
	}
	images := parse.GetElementsByTag(el, "img")
	for _, img := range images {
		src, ok := parse.FindAttr(img, "src")
		if ok {
			res = append(res, src)
		}
	}
	return res
}

func parseTableData(n *html.Node, label string) string {
	el := parseTableElement(n, label)
	if el == nil {
		return ""
	}
	return parse.GetText(el)
}

func parseLink(n *html.Node, label string) string {
	el := parseTableElement(n, label)
	if el == nil {
		return ""
	}
	a := parse.GetElementByTag(el, "a")
	if a == nil {
		return ""
	}
	href, ok := parse.FindAttr(a, "href")
	if !ok {
		return ""
	}
	match := detailLinkRegex.FindAllStringSubmatch(href, -1)
	if len(match) < 1 {
		return ""
	}
	return match[0][1]
}

func parseTableElement(n *html.Node, label string) *html.Node {
	info := parse.GetElementByClass(n, "inforbar_txt")
	if info == nil {
		return nil
	}
	tbody := parse.GetElementByTag(info, "tbody")
	if tbody == nil {
		return nil
	}
	rows := parse.GetElementsByTag(tbody, "tr")
	for _, row := range rows {
		elements := parse.GetElementsByTag(row, "td")
		if len(elements) == 2 && parse.GetText(elements[0]) == label {
			return elements[1]
		}
	}
	return nil
}
