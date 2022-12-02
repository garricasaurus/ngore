package internal

import (
	"git.okki.hu/garric/ngore/login"
	"git.okki.hu/garric/ngore/search"
	"net/url"
	"strconv"
)

func AuthForm(a login.Auth) url.Values {
	return url.Values{
		"nev":  {a.User()},
		"pass": {a.Pass()},
	}
}

func SearchForm(s *search.Params) url.Values {
	return url.Values{
		"mire":      {s.SearchPhrase},
		"miben":     {s.Field.String()},
		"tipus":     {s.Category.String()},
		"miszerint": {s.SortField.String()},
		"hogyan":    {s.SortMode.String()},
		"oldal":     {strconv.Itoa(s.Page + 1)},
	}
}
