package internal

import (
	"git.okki.hu/garric/ngore/search"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchForm(t *testing.T) {

	t.Run("page number indexing", func(t *testing.T) {
		s := &search.Params{
			Page: 0,
		}
		f := SearchForm(s)
		assert.Equal(t, []string{"1"}, f["oldal"])
	})

}
