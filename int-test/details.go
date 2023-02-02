package apitest

import "git.okki.hu/garric/ngore"

func PrintDetails(api ngore.Api) error {
	details, err := api.Details("3372123")
	if err != nil {
		return err
	}
	PrintJson(details)
	return nil
}
