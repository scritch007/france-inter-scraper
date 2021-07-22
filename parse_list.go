package fiscrap

import (
	"fmt"
	"golang.org/x/net/html"
)

func (s *Scrapper) Parse(url string) (map[string]string, error) {
	res := map[string]string{}

	l, err := s.parseList(url)
	if err != nil {
		return nil, err
	}

	for _, e := range l {
		toDownload, name, err := s.parsePodcastPage(e)
		if err != nil {
			fmt.Printf("Error retrieving podcast for %s\n", e)
			continue
		}
		res[name] = toDownload
	}

	return res, nil
}

func hasAttr(node *html.Node, attr, value string) bool {
	return getAttr(node, attr) == value
}

func getAttr(node *html.Node, attr string) string {
	for _, a := range node.Attr {
		if a.Key == attr {
			return a.Val
		}
	}
	return ""
}

func (s *Scrapper) parseList(url string) ([]string, error) {
	res := []string{}
	body, err := s.fetch.Do(url)
	if err != nil {
		return nil, err
	}

	n, err := html.Parse(body)
	if err != nil {
		return nil, err
	}
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {

		if node.Type == html.ElementNode && node.Data == "li" && hasAttr(node, "class", "tile") {

			res = append(res, parseTileUrl(node))
			return
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			crawler(child)
		}
	}

	crawler(n)

	return res, nil
}

func parseTileUrl(node *html.Node) string {
	v := getAttr(node.FirstChild.NextSibling.FirstChild.NextSibling.FirstChild, "href")
	return v
}
