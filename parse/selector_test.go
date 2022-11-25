package parse

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func TestGetElementsByClass(t *testing.T) {

	t.Run("no matching nodes", func(t *testing.T) {
		doc := mustParse(t, `<div><p>Test</p></div>`)
		nodes := GetElementsByClass(doc, "foo")
		assert.Equal(t, 0, len(nodes))
	})

	t.Run("multiple matches", func(t *testing.T) {
		doc := mustParse(t, `
		<div>
			<div id="elem1" class="foo" />
			<div class="bar" />
			<div>
				<p id="elem2" class="foo">Text</p>
			</div>
		</div>`)

		nodes := GetElementsByClass(doc, "foo")

		assert.Equal(t, 2, len(nodes))
		assert.Equal(t, "elem1", getId(nodes[0]))
		assert.Equal(t, "elem2", getId(nodes[1]))
	})

}

func TestGetElementById(t *testing.T) {

	doc := mustParse(t, `
		<div id="root">
			<a id="one" href="#" />
			<a id="two" href="#" />
		</div>`)

	t.Run("find element by id", func(t *testing.T) {
		n := GetElementById(doc, "two")
		assert.Equal(t, "two", getId(n))
	})

	t.Run("no element with matching id", func(t *testing.T) {
		n := GetElementById(doc, "foo")
		assert.Nil(t, n)
	})

}

func TestGetFirstChild(t *testing.T) {

	doc := mustParse(t, `
		<div id="root">
			<a id="one" href="#" />
			<a id="two" href="#" />
		</div>`)
	root := GetElementById(doc, "root")

	t.Run("find first matching child node", func(t *testing.T) {
		n := GetFirstChildWithTagName(root, "a")
		assert.Equal(t, "one", getId(n))
	})

	t.Run("no child nodes", func(t *testing.T) {
		doc := mustParse(t, `<div id="root" />`)
		root := GetElementById(doc, "root")
		assert.Nil(t, GetFirstChildWithTagName(root, "a"))
	})

	t.Run("no matching child node", func(t *testing.T) {
		assert.Nil(t, GetFirstChildWithTagName(root, "p"))
	})

}

func mustParse(t *testing.T, s string) *html.Node {
	t.Helper()
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		t.Fatal(err)
	}
	return doc
}

func getId(n *html.Node) string {
	for _, attr := range n.Attr {
		if attr.Key == "id" {
			return attr.Val
		}
	}
	return ""
}
