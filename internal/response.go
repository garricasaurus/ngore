package internal

import (
	"net/http"
	"strings"
)

func IsLoginRequired(res *http.Response) bool {
	return IsRedirectToLocation(res, LocationLogin)
}

func IsSuccessfulLogin(res *http.Response) bool {
	return IsRedirectToLocation(res, LocationIndex)
}

func IsInvalidLogin(res *http.Response) bool {
	return IsRedirectToLocation(res, LocationLoginProblem)
}

func IsRedirectToLocation(res *http.Response, location string) bool {
	if !IsRedirect(res) {
		return false
	}
	loc, _ := res.Location()
	return strings.Contains(loc.String(), location)
}

func IsRedirect(res *http.Response) bool {
	return res.StatusCode == http.StatusFound
}
