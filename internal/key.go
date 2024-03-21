package internal

import (
	"errors"
	"github.com/gar-r/ngore/parse"
	"golang.org/x/net/html"
	"strings"
)

func ExtractKey(doc *html.Node) (string, error) {
	links := parse.GetElementsByTag(doc, "link")
	for _, link := range links {
		href, ok := parse.FindAttr(link, "href")
		if ok && isRssRef(href) {
			return extractKey(href)
		}
	}
	return "", errors.New(ErrLoginKeyMissing)
}

func isRssRef(s string) bool {
	return strings.Contains(s, "rss.php")
}

func extractKey(s string) (string, error) {
	parts := strings.Split(s, "=")
	if len(parts) != 2 {
		return "", errors.New(ErrLoginKeyParse)
	}
	return parts[1], nil
}
