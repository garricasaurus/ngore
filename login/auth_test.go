package login

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicAuth_User(t *testing.T) {
	b := &BasicAuth{
		UserName: "user",
	}
	assert.Equal(t, "user", b.User())
}

func TestBasicAuth_Pass(t *testing.T) {
	b := &BasicAuth{
		Password: "pass",
	}
	assert.Equal(t, "pass", b.Pass())
}
