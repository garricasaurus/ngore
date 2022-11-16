package main

import (
	ncgo "git.okki.hu/garric/nc-go"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const EnvUser = "NC_USER"
const EnvPass = "NC_PASS"

const BaseUrl = "https://ncore.pro"

func TestUserJourney(t *testing.T) {

	api := ncgo.NewDefaultApi(BaseUrl, getAuthFromEnv())

	err := api.Login()
	assert.Nil(t, err)

}

func getAuthFromEnv() *ncgo.Credentials {
	return &ncgo.Credentials{
		User: os.Getenv(EnvUser),
		Pass: os.Getenv(EnvPass),
	}
}
