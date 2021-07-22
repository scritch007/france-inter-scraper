package fiscrap

import "io"

type Scrapper struct {
	fetch Fetcher
}

type Fetcher interface {
	Do(url string) (io.ReadCloser, error)
}

func New() *Scrapper {
	return &Scrapper{
		fetch: newFetcher(),
	}
}
