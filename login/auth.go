package login

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
