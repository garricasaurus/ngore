package main

import (
	"fmt"
	"git.okki.hu/garric/ngore"
	"git.okki.hu/garric/ngore/search"
)

func SearchMovieByName(api ngore.Api) error {
	params := &search.Params{
		SearchPhrase: "hortobagy",
		Field:        search.Name,
		Category:     search.MovieSdHu,
	}

	fmt.Printf("searching for '%s'\n", params.SearchPhrase)
	res, err := api.Search(params)
	if err != nil {
		return err
	}
	PrintJson(res)
	return nil
}
