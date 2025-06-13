package httpclient

import (
	"net/http"
	"net/url"
	"time"
)

// ClientManager manages HTTP clients and their configurations.
type ClientManager struct {
	defaultUserAgent string
}

// NewClientManager creates a new ClientManager.
func NewClientManager(defaultUserAgent string) *ClientManager {
	return &ClientManager{
		defaultUserAgent: defaultUserAgent,
	}
}

// CreateClient creates an HTTP client with specified timeout and optional proxy.
func (cm *ClientManager) CreateClient(timeout time.Duration, proxyURL *url.URL) *http.Client {
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	return &http.Client{
		Timeout:   timeout,
		Transport: transport,
	}
}

// CreateRequest creates an HTTP request with specified method, URL, and headers.
func (cm *ClientManager) CreateRequest(method, reqURL, cookie, userAgent string, httpVersion float64) (*http.Request, error) {
	req, err := http.NewRequest(method, reqURL, nil)
	if err != nil {
		return nil, err
	}

	if userAgent == "" {
		req.Header.Set("User-Agent", cm.defaultUserAgent)
	} else {
		req.Header.Set("User-Agent", userAgent)
	}

	if cookie != "" {
		req.Header.Set("Cookie", cookie)
	}

	if httpVersion == 2.0 {
		req.Proto = "HTTP/2.0"
	} else {
		req.Proto = "HTTP/1.1"
	}

	return req, nil
}


