package main

import (
	"git.okki.hu/garric/ncgore"
	"os"
)

const EnvUser = "NC_USER"
const EnvPass = "NC_PASS"

const BaseUrl = "https://ncore.pro"

func getAuthFromEnv() ncgore.Auth {
	return &ncgore.BasicAuth{
		UserName: os.Getenv(EnvUser),
		Password: os.Getenv(EnvPass),
	}
}
