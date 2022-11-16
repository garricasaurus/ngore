package ncgo

import (
	"net/http"
	"strings"
)

func isLoggedOut(res *http.Response) bool {
	return isRedirectToLocation(res, LocationLogin)
}

func isSuccessfulLogin(res *http.Response) bool {
	return isRedirectToLocation(res, LocationIndex)
}

func isInvalidLogin(res *http.Response) bool {
	return isRedirectToLocation(res, LocationLoginProblem)
}

func isRedirectToLocation(res *http.Response, location string) bool {
	if !isRedirect(res) {
		return false
	}
	loc, _ := res.Location()
	return strings.Contains(loc.String(), location)
}

func isRedirect(res *http.Response) bool {
	return res.StatusCode == http.StatusFound
}
