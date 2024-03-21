package internal

import (
	"net/url"
	"strconv"

	"github.com/gar-r/ngore/login"
	"github.com/gar-r/ngore/search"
)

func AuthForm(a login.Auth) url.Values {
	return url.Values{
		"nev":  {a.User()},
		"pass": {a.Pass()},
	}
}

func SearchForm(s *search.Params) url.Values {
	if s.Page < 1 {
		s.Page = 1
	}
	val := url.Values{}
	val.Set("mire", s.SearchPhrase)
	val.Set("miben", s.Field.String())
	val.Set("tipus", s.Category.String())
	val.Set("oldal", strconv.Itoa(s.Page))

	// do not apply sorting by name when we are searching by description
	// this is due to a bug in the website
	if s.Field == search.Description && s.SortField == search.ByName {
		return val
	}
	val.Set("miszerint", s.SortField.String())
	val.Set("hogyan", s.SortMode.String())
	return val
}
