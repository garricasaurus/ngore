package ncgo

import (
	"net/url"
)

type Credentials struct {
	User string
	Pass string
}

func (c *Credentials) Username() string {
	return c.User
}

func (c *Credentials) Password() string {
	return c.Pass
}

func (c *Credentials) ToFormValues() url.Values {
	return url.Values{
		"nev":  {c.Username()},
		"pass": {c.Password()},
	}
}
