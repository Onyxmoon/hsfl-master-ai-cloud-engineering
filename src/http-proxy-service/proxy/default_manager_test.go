package proxy

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewDefaultManager(t *testing.T) {
	config := &Config{
		ListenAddress: "localhost:8080",
		ProxyRoutes: []Route{
			{
				Name:    "TestRoute",
				Context: "/test",
				Target:  "http://example.com",
			},
		},
	}

	manager := NewDefaultManager(config)

	if manager == nil {
		t.Error("Expected a non-nil manager, got nil")
	}
}

func TestDefaultManager_GetProxyRouter(t *testing.T) {
	config := &Config{
		ListenAddress: "localhost:8080",
		ProxyRoutes: []Route{
			{
				Name:    "TestRoute",
				Context: "/test",
				Target:  "http://example.com",
			},
		},
	}
	manager := NewDefaultManager(config)

	router := manager.GetProxyRouter()

	if router == nil {
		t.Error("Expected a non-nil router, got nil")
	}
}

func TestProxyServer(t *testing.T) {
	want1 := "I'm service 1"
	want2 := "I'm service 2"

	testService1 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(want1))
	}))
	defer testService1.Close()
	testService2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(want2))
	}))
	defer testService2.Close()

	sampleConfig := &Config{
		ListenAddress: "localhost:8080",
		ProxyRoutes: []Route{
			{
				Name:    "Route1",
				Context: "/context1",
				Target:  testService1.URL,
			},
			{
				Name:    "Route2",
				Context: "/context2",
				Target:  testService2.URL,
			},
		},
	}

	proxyManager := NewDefaultManager(sampleConfig)
	testProxyServer := httptest.NewServer(proxyManager.GetProxyRouter())
	defer testProxyServer.Close()

	// Test an HTTP request to Route1
	t.Run("HTTP request to Route1", func(t *testing.T) {
		resp, err := testProxyServer.Client().Get(testProxyServer.URL + "/context1/")
		if err != nil {
			t.Errorf("HTTP GET to /context1/test failed: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		// Check response body matches
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Failed to read response body: %v", err)
		}
		if string(body) != want1 {
			t.Errorf("Expected response body for Route1 to be %s, but got %s", want1, string(body))
		}
	})

	t.Run("HTTP request to Route2", func(t *testing.T) {
		resp, err := testProxyServer.Client().Get(testProxyServer.URL + "/context2/")
		if err != nil {
			t.Errorf("HTTP GET to /context2/test failed: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		// Check response body matches
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Failed to read response body: %v", err)
		}
		if string(body) != want2 {
			t.Errorf("Expected response body for Route2 to be %s, but got %s", want2, string(body))
		}
	})
}
