package apitest

import (
	"errors"
	"fmt"
	"github.com/gar-r/ngore"
	"github.com/gar-r/ngore/search"
)

func Download(api ngore.Api) error {
	res, err := api.Search(&search.Params{
		SearchPhrase: "vadlovak hortobagy",
		Category:     search.MovieSdHu,
	})
	if err != nil {
		return err
	}
	if len(res.Torrents) == 0 {
		return errors.New("torrent not found")
	}
	t := res.Torrents[0]
	fmt.Printf("downloading %v\n", t)

	bytes, err := api.Download(t.Id)
	if err != nil {
		return err
	}

	fmt.Printf("download finished, size: %d\n", len(bytes))
	return nil
}
