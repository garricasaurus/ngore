package main

import (
	"git.okki.hu/garric/ngore"
	"git.okki.hu/garric/ngore/search"
	"testing"
)

func TestSearch(t *testing.T) {

	api := ngore.Default(BaseUrl)
	err := api.Login(getAuthFromEnv())
	if err != nil {
		t.Fatalf("error during login: %s", err)
	}

	t.Run("search sd movies by name", func(t *testing.T) {

		params := &search.Params{
			SearchPhrase: "hortobagy",
			Field:        search.Name,
			Category:     search.MovieSdHu,
		}

		t.Logf("search for: %v", params)

		res, err := api.Search(params)
		if err != nil {
			t.Errorf("search error: %s", err)
		}

		logResults(t, res)

	})

}

func logResults(t *testing.T, res []*search.Result) {
	for i, r := range res {
		t.Logf("[%d]: %v", i, r)
	}
}
