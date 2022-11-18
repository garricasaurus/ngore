package ncgore

import (
	"net/http"
	"net/http/cookiejar"
	"time"
)

type Api interface {
	Login(auth Auth) error
	Search(params *SearchParams) ([]*SearchResult, error)
}

type api struct {
	baseUrl string
	client  *http.Client
}

func New(client *http.Client, baseUrl string) Api {
	initCookieJar(client)
	disableRedirect(client)
	return &api{
		client:  client,
		baseUrl: baseUrl,
	}
}

func Default(baseUrl string) Api {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	return New(client, baseUrl)
}

func initCookieJar(client *http.Client) {
	jar, _ := cookiejar.New(nil)
	client.Jar = jar
}

func disableRedirect(client *http.Client) {
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
}
