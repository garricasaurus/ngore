package recommended

import (
	"regexp"

	"git.okki.hu/garric/ngore/parse"
	"golang.org/x/net/html"
)

var torrentRegex = regexp.MustCompile(`torrents\.php\?action=details&id=(\d+)`)

func ParseRecommendations(doc *html.Node) *Recommendations {
	result := &Recommendations{}
	result.Movies = parseSection(doc, "film")
	result.Series = parseSection(doc, "sorozat")
	result.Games = parseSection(doc, "jatek")
	result.Books = parseSection(doc, "ebook")
	result.Apps = parseSection(doc, "program")
	result.Music = parseSection(doc, "zene")
	return result
}

func parseSection(doc *html.Node, section string) *RecommendationGroup {
	a := findSection(doc, section)
	if a == nil {
		return nil
	}
	div := nextDiv(a)
	if div == nil {
		return nil
	}
	return parseRecommendations(div)
}

func nextDiv(n *html.Node) *html.Node {
	s := n.NextSibling
	for {
		if s == nil {
			return nil
		}
		if s.Type == html.ElementNode && s.Data == "div" {
			return s
		}
		s = s.NextSibling
	}
}

func parseRecommendations(n *html.Node) *RecommendationGroup {
	result := &RecommendationGroup{}
	anchors := parse.GetElementsByTag(n, "a")
	for _, a := range anchors {
		id := parseId(a)
		if id == "" {
			continue
		}
		img := parse.GetElementByTag(a, "img")
		if img == nil {
			rec := parseSimpleRec(a, id)
			result.Active = append(result.Active, rec)
		} else {
			rec := parseComplexRec(img, id)
			result.Staff = append(result.Staff, rec)
		}
	}
	return result
}

func parseComplexRec(img *html.Node, id string) *RecommendationWithImage {
	rec := &RecommendationWithImage{}
	rec.Id = id
	title, _ := parse.FindAttr(img, "title")
	rec.Name = title
	src, _ := parse.FindAttr(img, "src")
	rec.Image = src
	return rec
}

func parseSimpleRec(a *html.Node, id string) *Recommendation {
	return &Recommendation{
		Id:   id,
		Name: parse.GetText(a),
	}
}

func parseId(n *html.Node) string {
	href, ok := parse.FindAttr(n, "href")
	if !ok {
		return ""
	}
	matches := torrentRegex.FindAllStringSubmatch(href, -1)
	if len(matches) < 1 {
		return ""
	}
	return matches[0][1]
}

func findSection(n *html.Node, name string) *html.Node {
	anchors := parse.GetElementsByTag(n, "a")
	for _, a := range anchors {
		val, ok := parse.FindAttr(a, "name")
		if ok && val == name {
			return a
		}
	}
	return nil
}
