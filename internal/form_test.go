package internal

import (
	"github.com/gar-r/ngore/login"
	"github.com/gar-r/ngore/search"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthForm(t *testing.T) {
	a := &login.BasicAuth{
		UserName: "user",
		Password: "pass",
	}
	val := AuthForm(a)
	assert.Equal(t, "user", val.Get("nev"))
	assert.Equal(t, "pass", val.Get("pass"))
}

func TestSearchForm(t *testing.T) {

	t.Run("page number indexing", func(t *testing.T) {
		s := &search.Params{
			Page: 0,
		}
		f := SearchForm(s)
		assert.Equal(t, []string{"1"}, f["oldal"])
	})

	t.Run("sort params", func(t *testing.T) {
		s := &search.Params{
			Field:     search.Name,
			SortField: search.ByName,
			SortMode:  search.Descending,
		}
		f := SearchForm(s)
		assert.Equal(t, search.ByName.String(), f.Get("miszerint"))
		assert.Equal(t, search.Descending.String(), f.Get("hogyan"))
	})

	t.Run("ncore search by description name sort bug", func(t *testing.T) {
		s := &search.Params{
			Field:     search.Description,
			SortField: search.ByName,
			SortMode:  search.Descending,
		}
		f := SearchForm(s)
		assert.False(t, f.Has("miszerint"))
		assert.False(t, f.Has("hogyan"))
	})

}
