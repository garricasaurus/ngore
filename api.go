package ngore

import (
	"errors"
	"fmt"
	"git.okki.hu/garric/ngore/internal"
	"git.okki.hu/garric/ngore/login"
	"git.okki.hu/garric/ngore/search"
	"golang.org/x/net/html"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type Api interface {
	Login(auth login.Auth) error
	Search(params *search.Params) (*search.Result, error)
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

func (a *api) Login(auth login.Auth) error {
	if auth.User() == "" || auth.Pass() == "" {
		return errors.New(internal.ErrLoginMissingCredentials)
	}
	res, err := a.client.PostForm(a.baseUrl+internal.UrlLogin, internal.AuthForm(auth))
	if err != nil {
		return err
	}
	if internal.IsInvalidLogin(res) {
		return errors.New(internal.ErrLoginInvalidCredentials)
	}
	if internal.IsSuccessfulLogin(res) {
		return nil

	}
	return errors.New(internal.ErrLoginUnexpectedResponse)
}

func (a *api) Search(params *search.Params) (*search.Result, error) {
	res, err := a.client.PostForm(a.baseUrl+internal.UrlSearch, internal.SearchForm(params))
	if err != nil {
		return nil, err
	}
	if internal.IsLoginRequired(res) {
		return nil, errors.New(internal.ErrUserNotLoggedIn)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(internal.ErrSearchUnexpectedResponseCode, res.StatusCode)
	}
	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil, err
	}
	return search.ParseResponse(doc), nil
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
