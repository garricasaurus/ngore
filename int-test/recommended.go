package apitest

import "github.com/gar-r/ngore"

func PrintRecommendations(api ngore.Api) error {
	rec, err := api.Recommendations()
	if err != nil {
		return err
	}
	PrintJson(rec)
	return nil
}
