package httpclient

import "net/http"

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

type DefaultHTTPClient struct{}

func NewDefaultHTTPClient() *DefaultHTTPClient {
	return &DefaultHTTPClient{}
}

func (c *DefaultHTTPClient) Get(url string) (*http.Response, error) {
	return http.Get(url)
}
