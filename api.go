package ncgo

import (
	"errors"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type Api struct {
	BaseUrl     string
	Client      *http.Client
	Credentials *Credentials
}

func NewApi(client *http.Client, baseUrl string, credentials *Credentials) *Api {
	api := &Api{
		Client:      client,
		BaseUrl:     baseUrl,
		Credentials: credentials,
	}
	api.disableRedirect()
	return api
}

func NewDefaultApi(baseUrl string, credentials *Credentials) *Api {
	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Timeout: 10 * time.Second,
		Jar:     jar,
	}
	return NewApi(client, baseUrl, credentials)
}

func (a *Api) Login() error {
	res, err := a.Client.PostForm(a.BaseUrl+UrlLogin, a.Credentials.ToFormValues())
	if err != nil {
		return err
	}
	if isInvalidLogin(res) {
		return errors.New(ErrLoginInvalidCredentials)
	}
	if isSuccessfulLogin(res) {
		return nil

	}
	return errors.New(ErrLoginUnexpectedResponse)
}

func (a *Api) disableRedirect() {
	a.Client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
}
