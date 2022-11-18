package ncgore

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewApi(t *testing.T) {

	t.Run("new api not nil", func(t *testing.T) {
		api := New(&http.Client{}, "baseUrl")
		assert.NotNil(t, api)
	})

	t.Run("redirect is disabled", func(t *testing.T) {
		client := &http.Client{}
		_ = New(client, "baseUrl")
		assert.Equal(t, http.ErrUseLastResponse, client.CheckRedirect(nil, nil))
	})

	t.Run("cookie jar initialized", func(t *testing.T) {
		api := New(&http.Client{}, "baseUrl").(*api)
		assert.NotNil(t, api.client.Jar)
	})

}

func TestNewDefaultApi(t *testing.T) {

	t.Run("client initialized", func(t *testing.T) {
		api := Default("baseUrl").(*api)
		assert.NotNil(t, api.client)
	})

	t.Run("default timeout", func(t *testing.T) {
		api := Default("baseUrl").(*api)
		assert.Greater(t, api.client.Timeout, time.Duration(0))
	})

	t.Run("default cookie jar", func(t *testing.T) {
		api := Default("baseUrl").(*api)
		assert.NotNil(t, api.client.Jar)
	})

}

func apiWithMockClient(mockServer *httptest.Server) Api {
	return New(mockServer.Client(), mockServer.URL)
}
