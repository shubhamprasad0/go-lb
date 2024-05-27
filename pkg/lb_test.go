package lb

import (
	"fmt"
	"net/http"
	"testing"

	testUtils "github.com/shubhamprasad0/go-lb/test"
	"github.com/stretchr/testify/assert"
)

func TestLoadBalancer(t *testing.T) {
	// Start application servers
	ts1 := testUtils.StartTestServer("hello from server 1", http.StatusOK)
	defer ts1.Close()
	ts2 := testUtils.StartTestServer("hello from server 2", http.StatusOK)
	defer ts2.Close()
	ts3 := testUtils.StartTestServer("hello from server 3", http.StatusOK)
	defer ts3.Close()

	config := DefaultLBConfig()
	config.Servers = []string{ts1.URL[7:], ts2.URL[7:], ts3.URL[7:]} // remove "http://"
	config.HealthCheckRoute = "/"
	lb := NewLoadBalancer(config)
	lb.healthChecker.performHealthChecks()

	go lb.Start()

	var responses []string
	for i := 0; i < 30; i++ {
		client := &http.Client{}
		resp, err := client.Get(fmt.Sprintf("http://localhost:%d", lb.Config.Port))
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		data := make([]byte, 1000)
		n, _ := resp.Body.Read(data)
		responseBody := string(data[:n])
		assert.Contains(t, responseBody, "hello from server")

		// store response from server
		responses = append(responses, responseBody)
	}

	// check even distribution
	count1 := 0
	count2 := 0
	count3 := 0
	for _, resp := range responses {
		if resp == "hello from server 1" {
			count1++
		} else if resp == "hello from server 2" {
			count2++
		} else if resp == "hello from server 3" {
			count3++
		}
	}
	assert.Equal(t, 10, count1)
	assert.Equal(t, 10, count2)
	assert.Equal(t, 10, count3)
}
