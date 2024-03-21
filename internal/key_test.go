package internal

import (
	"github.com/gar-r/ngore/parse"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractKey(t *testing.T) {

	t.Run("extract key", func(t *testing.T) {
		doc := parse.MustParse(t, `
		<html>
			<head>
			<link rel="alternate" href="something" title="bar">
			<link rel="alternate" href="/rss.php?key=secret_key" title="foo">
			</head>
		</html>`)
		key, err := ExtractKey(doc)
		assert.NoError(t, err)
		assert.Equal(t, "secret_key", key)
	})

	t.Run("no rss link", func(t *testing.T) {
		doc := parse.MustParse(t, `<link title="foo" />`)
		_, err := ExtractKey(doc)
		assert.Error(t, err)
	})

	t.Run("invalid rss link", func(t *testing.T) {
		doc := parse.MustParse(t, `<link href="/rss.php?foo=bar&key=secret_key" title="foo">`)
		_, err := ExtractKey(doc)
		assert.Error(t, err)
	})

}
