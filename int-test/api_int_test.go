package apitest

import (
	"os"
	"testing"

	"github.com/gar-r/ngore"
	"github.com/gar-r/ngore/login"
	"github.com/stretchr/testify/assert"
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
		t.Run("search movie by name", func(t *testing.T) {
			assert.NoError(t, SearchMovieByName(api))
		})

		t.Run("search with paging", func(t *testing.T) {
			assert.NoError(t, SearchWithPaging(api))
		})

		t.Run("search by description", func(t *testing.T) {
			assert.NoError(t, SearchByDescription(api))
		})

		t.Run("search with sort", func(t *testing.T) {
			assert.NoError(t, SearchWithSort(api))
		})
	})

	t.Run("activity", func(t *testing.T) {
		assert.NoError(t, PrintActivity(api))
	})

	t.Run("download", func(t *testing.T) {
		assert.NoError(t, Download(api))
	})

	t.Run("details", func(t *testing.T) {
		assert.NoError(t, PrintDetails(api))
	})

	t.Run("recommendations", func(t *testing.T) {
		assert.NoError(t, PrintRecommendations(api))
	})

}
