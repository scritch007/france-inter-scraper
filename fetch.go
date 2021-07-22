package fiscrap

import (
	"io"
	"net/http"
)

type fetch struct{}

func newFetcher() *fetch {
	return &fetch{}
}

func (f *fetch) Do(url string) (io.ReadCloser, error) {
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
