package ncgore

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApi_Login(t *testing.T) {

	t.Run("successful login", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Set-Cookie", "PHPSESSID=foo")
			http.Redirect(w, r, "https://example.com/index.php", http.StatusFound)
		}))
		defer server.Close()
		api := apiWithMockClient(server)
		err := api.Login(&BasicAuth{"user", "pass"})
		assert.Nil(t, err)
	})

	t.Run("invalid login", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "https://example.com?problema=1", http.StatusFound)
		}))
		defer server.Close()
		api := apiWithMockClient(server)
		err := api.Login(&BasicAuth{"user", "pass"})
		assert.ErrorContains(t, err, errLoginInvalidCredentials)
	})

	t.Run("missing login details", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()
		api := apiWithMockClient(server)
		err := api.Login(&BasicAuth{})
		assert.ErrorContains(t, err, errLoginMissingCredentials)
	})

	t.Run("server error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "test error", http.StatusInternalServerError)
		}))
		defer server.Close()
		api := apiWithMockClient(server)
		err := api.Login(&BasicAuth{"user", "pass"})
		assert.ErrorContains(t, err, errLoginUnexpectedResponse)
	})

	t.Run("other error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
		api := apiWithMockClient(server)
		server.Close() // close server to cause an error in login
		err := api.Login(&BasicAuth{"user", "pass"})
		assert.Error(t, err)
	})

}
