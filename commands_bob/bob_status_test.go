package commands_bob

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReturnBobStatusCode(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.RequestURI == "/api/status" {
			_, _ = rw.Write([]byte(`{"message":"Running"}`))
		}
	}))

	bobResponse := BobResponse{Message:"Running"}

	api := API{
		Client: server.Client(),
		BaseURL: server.URL,
	}

	// Close the server when test finishes
	defer server.Close()
	status, _ := api.RunningStatus()
	assert.Equal(t, bobResponse, status)
}
