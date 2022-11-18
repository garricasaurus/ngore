package ncgore

import (
	"errors"
	"net/url"
)

func (a *api) Login(auth Auth) error {
	if blank(auth) {
		return errors.New(errLoginMissingCredentials)
	}
	res, err := a.client.PostForm(a.baseUrl+urlLogin, authForm(auth))
	if err != nil {
		return err
	}
	if isInvalidLogin(res) {
		return errors.New(errLoginInvalidCredentials)
	}
	if isSuccessfulLogin(res) {
		return nil

	}
	return errors.New(errLoginUnexpectedResponse)
}

type Auth interface {
	User() string
	Pass() string
}

type BasicAuth struct {
	UserName string
	Password string
}

func (b *BasicAuth) User() string {
	return b.UserName
}

func (b *BasicAuth) Pass() string {
	return b.Password
}

func authForm(a Auth) url.Values {
	return url.Values{
		"nev":  {a.User()},
		"pass": {a.Pass()},
	}
}

func blank(a Auth) bool {
	return a.User() == "" || a.Pass() == ""
}
