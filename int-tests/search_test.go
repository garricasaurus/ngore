package main

import (
	"git.okki.hu/garric/ncgore"
	"testing"
)

func TestSearch(t *testing.T) {

	api := ncgore.Default(BaseUrl)
	err := api.Login(getAuthFromEnv())
	if err != nil {
		t.Fatalf("error during login: %s", err)
	}

	t.Run("search sd movies by name", func(t *testing.T) {

		params := &ncgore.SearchParams{
			SearchPhrase: "hortobagy",
			Field:        ncgore.Name,
			Category:     ncgore.MovieSdHu,
		}

		t.Logf("search for: %v", params)

		res, err := api.Search(params)
		if err != nil {
			t.Errorf("search error: %s", err)
		}

		logResults(t, res)

	})

}

func logResults(t *testing.T, res []*ncgore.SearchResult) {
	for i, r := range res {
		t.Logf("[%d]: %v", i, r)
	}
}
