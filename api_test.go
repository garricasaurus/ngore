package ncgo

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewApi(t *testing.T) {

	t.Run("new api not nil", func(t *testing.T) {
		api := NewApi(&http.Client{}, "baseUrl", &Credentials{})
		assert.NotNil(t, api)
	})

	t.Run("redirect is disabled", func(t *testing.T) {
		client := &http.Client{}
		_ = NewApi(client, "baseUrl", &Credentials{})
		assert.Equal(t, http.ErrUseLastResponse, client.CheckRedirect(nil, nil))
	})

}

func TestNewDefaultApi(t *testing.T) {

	t.Run("default timeout", func(t *testing.T) {
		api := NewDefaultApi("baseUrl", &Credentials{})
		assert.Greater(t, api.Client.Timeout, time.Duration(0))
	})

	t.Run("default cookie jar", func(t *testing.T) {
		api := NewDefaultApi("baseUrl", &Credentials{})
		assert.NotNil(t, api.Client.Jar)
	})

}

func TestApi_Login(t *testing.T) {

	t.Run("successful login", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Set-Cookie", "PHPSESSID=foo")
			http.Redirect(w, r, "https://example.com/index.php", http.StatusFound)
		}))
		defer server.Close()
		api := apiWithMockClient(server)
		err := api.Login()
		assert.Nil(t, err)
	})

	t.Run("invalid credentials", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://example.com?problema=1", http.StatusFound)
		}))
		defer server.Close()
		api := apiWithMockClient(server)
		err := api.Login()
		assert.ErrorContains(t, err, ErrLoginInvalidCredentials)
	})

	t.Run("server error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "test error", http.StatusInternalServerError)
		}))
		defer server.Close()
		api := apiWithMockClient(server)
		err := api.Login()
		assert.ErrorContains(t, err, ErrLoginUnexpectedResponse)
	})

	t.Run("other error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		api := apiWithMockClient(server)
		server.Close() // close server to cause an error in login
		err := api.Login()
		assert.Error(t, err)
	})

}

func apiWithMockClient(mockServer *httptest.Server) *Api {
	return NewApi(mockServer.Client(), mockServer.URL, &Credentials{})
}
