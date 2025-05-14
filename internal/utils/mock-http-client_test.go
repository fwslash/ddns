package utils

import (
	"io"
	"net/http"
	"strings"
)

type MockClient struct {
	ResponseBody string
	Err          error
}

func (m *MockClient) Get(url string) (*http.Response, error) {
	if m.Err != nil {
		return nil, m.Err
	}

	body := io.NopCloser(strings.NewReader(m.ResponseBody))
	return &http.Response{
		StatusCode: 200,
		Body:       body,
	}, nil
}
