package apitest

import (
	"github.com/gar-r/ngore"
)

func PrintActivity(api ngore.Api) error {
	info, err := api.Activity()
	if err != nil {
		return err
	}
	PrintJson(info)
	return nil
}
