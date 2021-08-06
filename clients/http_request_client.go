package clients

import (
	"io"
	"net/http"
)

type HttpRequestClient interface {
	DoRequest(method string, url string, body io.Reader) (*http.Response, error)
}

type HttpRequestClientImpl struct {
	client *http.Client
}

func NewHttpRequestClient(client *http.Client) HttpRequestClient {
	return &HttpRequestClientImpl{client: client}
}

func (h *HttpRequestClientImpl) DoRequest(method string, url string, body io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	return h.client.Do(request)
}
