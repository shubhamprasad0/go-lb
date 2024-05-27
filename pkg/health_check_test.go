package lb

import (
	"net/http"
	"testing"

	testUtils "github.com/shubhamprasad0/go-lb/test"
	"github.com/stretchr/testify/assert"
)

func TestHealthChecker(t *testing.T) {

	ts1 := testUtils.StartTestServer("", http.StatusOK)
	defer ts1.Close()
	ts2 := testUtils.StartTestServer("", http.StatusInternalServerError)
	defer ts2.Close()

	servers := []string{ts1.URL[7:], ts2.URL[7:]} // remove "http://"
	healthChecker := NewHealthChecker(servers, "/", 5)

	healthChecker.performHealthChecks()

	assert.Contains(t, healthChecker.HealthyServers, ts1.URL[7:]) // remove "http://"
	assert.NotContains(t, healthChecker.HealthyServers, ts2.URL[7:])
}
