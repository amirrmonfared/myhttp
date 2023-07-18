package httpclient

import "net/http"

type FakeHTTPClient struct {
	GetResponse *http.Response
	Err         error
}

func (c *FakeHTTPClient) Get(url string) (*http.Response, error) {
	return c.GetResponse, c.Err
}
