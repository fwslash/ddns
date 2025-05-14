package utils


import(
  "net/http"
  "time"
)


type HttpClient interface {
	Get(url string) (*http.Response, error)
}

func Client() HttpClient {
  return &http.Client{Timeout: 30 * time.Second}
}

