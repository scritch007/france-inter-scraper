package fiscrap

import (
	"fmt"
	"golang.org/x/net/html"
	url2 "net/url"
)

func (s *Scrapper) parsePodcastPage(url string) (string, string, error) {
	u, err := url2.Parse(url)
	if err != nil {
		return "", "", err
	}

	p, err := s.fetch.Do(url)
	if err != nil {
		return "", "", err
	}
	n, err := html.Parse(p)
	if err != nil {
		return "", "", err
	}
	var crawler func(*html.Node) (string, string, bool)

	crawler = func(node *html.Node) (string, string, bool) {
		if node.Type == html.ElementNode && node.Data == "button"{
			if hasAttr(node, "data-diffusion-path", u.Path) {
				return getAttr(node, "data-url"), getAttr(node, "title"), true
			}
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			a, b, f := crawler(child)
			if f {
				return a, b, f
			}
		}
		return "", "", false
	}
	a, b, f := crawler(n)
	if !f {
		return "", "", fmt.Errorf("link not found")
	}
	return a, b, nil

}
