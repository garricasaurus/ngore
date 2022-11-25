package main

import (
	"git.okki.hu/garric/ncgore/login"
	"os"
)

const EnvUser = "NC_USER"
const EnvPass = "NC_PASS"

const BaseUrl = "https://ncore.pro"

func getAuthFromEnv() login.Auth {
	return &login.BasicAuth{
		UserName: os.Getenv(EnvUser),
		Password: os.Getenv(EnvPass),
	}
}
