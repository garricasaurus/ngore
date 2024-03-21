package ngore

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/gar-r/ngore/activity"
	"github.com/gar-r/ngore/details"
	"github.com/gar-r/ngore/internal"
	"github.com/gar-r/ngore/login"
	"github.com/gar-r/ngore/recommended"
	"github.com/gar-r/ngore/search"
	"golang.org/x/net/html"
)

type Api interface {
	Login(auth login.Auth) error
	Search(params *search.Params) (*search.Result, error)
	Activity() (*activity.Info, error)
	Recommendations() (*recommended.Recommendations, error)
	Details(id string) (*details.Details, error)
	Download(id string) ([]byte, error)
}

type api struct {
	baseUrl string
	key     string
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
		return a.fetchKey()
	}
	return errors.New(internal.ErrLoginUnexpectedResponse)
}

func (a *api) Search(params *search.Params) (*search.Result, error) {
	res, err := a.client.PostForm(a.baseUrl+internal.UrlTorrents, internal.SearchForm(params))
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

func (a *api) Activity() (*activity.Info, error) {
	res, err := a.client.Get(a.baseUrl + internal.UrlActivity)
	if err != nil {
		return nil, err
	}
	if internal.IsLoginRequired(res) {
		return nil, errors.New(internal.ErrUserNotLoggedIn)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(internal.ErrActivityUnexpectedResponseCode, res.StatusCode)
	}
	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil, err
	}
	return activity.ParseResponse(doc), nil
}

func (a *api) Recommendations() (*recommended.Recommendations, error) {
	res, err := a.client.Get(a.baseUrl + internal.UrlRecommended)
	if err != nil {
		return nil, err
	}
	if internal.IsLoginRequired(res) {
		return nil, errors.New(internal.ErrUserNotLoggedIn)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(internal.ErrActivityUnexpectedResponseCode, res.StatusCode)
	}
	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil, err
	}
	return recommended.ParseRecommendations(doc), nil
}

func (a *api) Download(id string) ([]byte, error) {
	if a.key == "" {
		return nil, errors.New(internal.ErrApiKeyEmpty)
	}
	query := fmt.Sprintf("?action=download&id=%s&key=%s", id, a.key)
	url := a.baseUrl + internal.UrlTorrents + query
	res, err := a.client.Get(url)
	if err != nil {
		return nil, err
	}
	if internal.IsLoginRequired(res) {
		return nil, errors.New(internal.ErrUserNotLoggedIn)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(internal.ErrDownloadUnexpectedResponseCode, res.StatusCode)
	}
	return io.ReadAll(res.Body)
}

func (a *api) Details(id string) (*details.Details, error) {
	query := fmt.Sprintf("?action=details&id=%s", id)
	url := a.baseUrl + internal.UrlTorrents + query
	res, err := a.client.Get(url)
	if err != nil {
		return nil, err
	}
	if internal.IsLoginRequired(res) {
		return nil, errors.New(internal.ErrUserNotLoggedIn)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(internal.ErrDetailsUnexpectedResponseCode, res.StatusCode)
	}
	doc, err := html.Parse(res.Body)
	if err != nil {
		return nil, err
	}
	return details.ParseDetails(doc), nil
}

func (a *api) fetchKey() error {
	res, err := a.client.Get(a.baseUrl + internal.UrlIndex)
	if err != nil {
		return errors.New(internal.ErrLoginUnableToFetchIndex)
	}
	doc, err := html.Parse(res.Body)
	if err != nil {
		return err
	}
	a.key, err = internal.ExtractKey(doc)
	return err
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
