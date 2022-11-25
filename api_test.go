package ncgore

import (
	"git.okki.hu/garric/ncgore/internal"
	"git.okki.hu/garric/ncgore/login"
	"git.okki.hu/garric/ncgore/search"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewApi(t *testing.T) {

	t.Run("new api not nil", func(t *testing.T) {
		ng := New(&http.Client{}, "baseUrl")
		assert.NotNil(t, ng)
	})

	t.Run("redirect is disabled", func(t *testing.T) {
		client := &http.Client{}
		_ = New(client, "baseUrl")
		assert.Equal(t, http.ErrUseLastResponse, client.CheckRedirect(nil, nil))
	})

	t.Run("cookie jar initialized", func(t *testing.T) {
		ng := New(&http.Client{}, "baseUrl").(*api)
		assert.NotNil(t, ng.client.Jar)
	})

}

func TestNewDefaultApi(t *testing.T) {

	t.Run("client initialized", func(t *testing.T) {
		ng := Default("baseUrl").(*api)
		assert.NotNil(t, ng.client)
	})

	t.Run("default timeout", func(t *testing.T) {
		ng := Default("baseUrl").(*api)
		assert.Greater(t, ng.client.Timeout, time.Duration(0))
	})

	t.Run("default cookie jar", func(t *testing.T) {
		ng := Default("baseUrl").(*api)
		assert.NotNil(t, ng.client.Jar)
	})

}

func TestApi_Login(t *testing.T) {

	t.Run("successful login", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Set-Cookie", "PHPSESSID=foo")
			http.Redirect(w, r, "https://example.com/index.php", http.StatusFound)
		}))
		defer server.Close()
		ng := apiWithMockClient(server)
		err := ng.Login(&login.BasicAuth{UserName: "user", Password: "pass"})
		assert.Nil(t, err)
	})

	t.Run("invalid login", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://example.com?problema=1", http.StatusFound)
		}))
		defer server.Close()
		ng := apiWithMockClient(server)
		err := ng.Login(&login.BasicAuth{UserName: "user", Password: "pass"})
		assert.ErrorContains(t, err, internal.ErrLoginInvalidCredentials)
	})

	t.Run("server error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "test error", http.StatusInternalServerError)
		}))
		defer server.Close()
		ng := apiWithMockClient(server)
		err := ng.Login(&login.BasicAuth{UserName: "user", Password: "pass"})
		assert.ErrorContains(t, err, internal.ErrLoginUnexpectedResponse)
	})

	t.Run("other error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		ng := apiWithMockClient(server)
		server.Close() // close server to cause an error in login
		err := ng.Login(&login.BasicAuth{UserName: "user", Password: "pass"})
		assert.Error(t, err)
	})

}

func TestApi_Search(t *testing.T) {

	t.Run("search api login required", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, internal.LocationLogin, http.StatusFound)
		}))
		defer server.Close()
		api := apiWithMockClient(server)
		_, err := api.Search(&search.Params{})
		assert.ErrorContains(t, err, internal.ErrUserNotLoggedIn)
	})

	t.Run("search api server error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()
		api := apiWithMockClient(server)
		_, err := api.Search(&search.Params{})
		assert.ErrorContains(t, err, "500")
	})

	t.Run("search api network error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		api := apiWithMockClient(server)
		server.Close() // close server to cause an error
		_, err := api.Search(&search.Params{})
		assert.Error(t, err)
	})

}

func apiWithMockClient(mockServer *httptest.Server) Api {
	return New(mockServer.Client(), mockServer.URL)
}
