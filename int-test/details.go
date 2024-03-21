package apitest

import "github.com/gar-r/ngore"

func PrintDetails(api ngore.Api) error {
	details, err := api.Details("3372123")
	if err != nil {
		return err
	}
	PrintJson(details)
	return nil
}
