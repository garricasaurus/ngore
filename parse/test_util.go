package parse

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func MustParse(t *testing.T, s string) *html.Node {
	t.Helper()
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		t.Fatal(err)
	}
	return doc
}
