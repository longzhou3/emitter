package security

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testVaultHandler struct{}

func (h *testVaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := ""
	switch r.URL.Path {
	case "/v1/auth/app-id/login":
		response = `{ "auth": {"client_token": "123" } }`
	case "/v1/secret/test":
		response = `{ "data": {"value": "hi" } }`
	case "/v1/aws/sts/storage-emitter":
		response = `{ "lease_duration": 100, "data": { "access_key": "access", "secret_key": "secret", "security_token": "token" } }`
	}

	w.Write([]byte(response))
}

func TestVaultClient(t *testing.T) {
	s := httptest.NewServer(&testVaultHandler{})
	defer s.Close()

	// Test authentication first
	cli := newVaultClient(s.URL)
	err := cli.Authenticate("xxxxxx", "yyyyyy")
	assert.NoError(t, err)
	assert.Equal(t, "123", cli.token)

	// Test secret read now
	{
		v, err := cli.ReadSecret("test")
		assert.NoError(t, err)
		assert.Equal(t, "hi", v)
	}

	// Test credentials read now
	{
		v, err := cli.ReadCredentials("storage-emitter")
		assert.NoError(t, err)
		assert.Equal(t, "access", v.AccessKey)
		assert.Equal(t, "secret", v.SecretKey)
		assert.Equal(t, "token", v.Token)
	}
}