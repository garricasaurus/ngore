package apitest

import (
	"git.okki.hu/garric/ngore"
	"git.okki.hu/garric/ngore/login"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

const EnvUser = "NC_USER"
const EnvPass = "NC_PASS"

const BaseUrl = "https://ncore.pro"

func TestApi(t *testing.T) {

	var api ngore.Api

	t.Run("create api", func(t *testing.T) {
		api = ngore.Default(BaseUrl)
		assert.NotNil(t, api)
	})

	t.Run("login", func(t *testing.T) {
		auth := &login.BasicAuth{
			UserName: os.Getenv(EnvUser),
			Password: os.Getenv(EnvPass),
		}
		assert.NoError(t, api.Login(auth))
	})

	t.Run("search", func(t *testing.T) {
		assert.NoError(t, SearchMovieByName(api))
		assert.NoError(t, SearchWithPaging(api))
	})

	t.Run("activity", func(t *testing.T) {
		assert.NoError(t, PrintActivity(api))
	})

}
