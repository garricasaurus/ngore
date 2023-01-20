package apitest

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

func SearchWithPaging(api ngore.Api) error {
	params := &search.Params{
		SearchPhrase: "game",
		Field:        search.Name,
		Category:     search.SeriesSdEn,
	}

	fmt.Printf("searching for %s\n", params.SearchPhrase)

	var res *search.Result
	var err error
	for params.Page = 0; res == nil || res.Page.HasMore(); params.Page++ {
		fmt.Printf("page %d\n", params.Page)
		res, err = api.Search(params)
		if err != nil {
			return err
		}
		PrintJson(res)
	}
	return nil
}

func SearchByDescription(api ngore.Api) error {
	params := &search.Params{
		SearchPhrase: "Arctic Secrets",
		Field:        search.Description,
		Category:     search.SeriesSdEn,
	}
	res, err := api.Search(params)
	if err != nil {
		return err
	}
	PrintJson(res)
	return nil
}
