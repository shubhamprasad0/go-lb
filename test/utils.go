package test_utils

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper functions
func CreateTempFile(t *testing.T, name string, content string) *os.File {
	file, err := os.CreateTemp("", name)
	assert.NoError(t, err)

	_, err = file.Write([]byte(content))
	assert.NoError(t, err)

	err = file.Close()
	assert.NoError(t, err)

	return file
}

func StartTestServer(data string, statusCode int) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		if data != "" {
			w.Write([]byte(data))
		}
	})
	ts := httptest.NewServer(handler)
	return ts
}
