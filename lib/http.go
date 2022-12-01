package lib

import (
	"io"
	"net/http"
)

type Requester interface {
	Get(url string) (*http.Response, error)
	Post(url string, body io.Reader) (*http.Response, error)
	Patch(url string, body io.Reader) (*http.Response, error)
	Delete(url string, body io.Reader) (*http.Response, error)
}

type HttpClient struct {
	client *http.Client
}

func NewHttpClient(client *http.Client) *HttpClient {
	return &HttpClient{
		client: client,
	}
}

func (c *HttpClient) Get(url string) (*http.Response, error) {
	return c.client.Get(url)
}

func (c *HttpClient) Post(url string, body io.Reader) (*http.Response, error) {
	return c.client.Post(url, "application/json", body)
}

func (c *HttpClient) Patch(url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("PATCH", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return c.client.Do(req)
}

func (c *HttpClient) Delete(url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", url, body)
	if err != nil {
		return nil, err
	}
	return c.client.Do(req)
}
