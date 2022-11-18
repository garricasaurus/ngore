package ncgore

import (
	"net/http"
	"strings"
)

func isLoginRequired(res *http.Response) bool {
	return isRedirectToLocation(res, locationLogin)
}

func isSuccessfulLogin(res *http.Response) bool {
	return isRedirectToLocation(res, locationIndex)
}

func isInvalidLogin(res *http.Response) bool {
	return isRedirectToLocation(res, locationLoginProblem)
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
