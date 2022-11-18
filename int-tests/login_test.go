package main

import (
	"git.okki.hu/garric/ncgore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogin(t *testing.T) {

	api := ncgore.Default(BaseUrl)

	err := api.Login(getAuthFromEnv())
	assert.Nil(t, err)

	t.Logf("login successful")
}
