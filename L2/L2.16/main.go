/* Реализовать утилиту загрузки веб-страниц вместе со всем вложенным контентом
(ресурсы, ссылки), аналогичную wget -m (мирроринг сайта). */

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"golang.org/x/net/html"
)

type Crawler struct {
	maxDepth int
	visited  map[string]bool
}

func main() {
	depth := flag.Int("depth", 2, "recursion depth")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("usage: wget [-depth N] URL")
		return
	}

	crawler := &Crawler{
		maxDepth: *depth,
		visited:  make(map[string]bool),
	}

	err := crawler.crawl(flag.Arg(0), 0)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func (c *Crawler) crawl(rawURL string, depth int) error {
	if depth > c.maxDepth {
		return nil
	}

	if c.visited[rawURL] {
		return nil
	}

	c.visited[rawURL] = true

	fmt.Println("Downloading:", rawURL)

	resp, err := http.Get(rawURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = saveFile(rawURL, body)
	if err != nil {
		return err
	}

	contentType := resp.Header.Get("Content-Type")

	if len(contentType) < 9 || contentType[:9] != "text/html" {
		return nil
	}

	doc, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		return err
	}

	baseURL, err := url.Parse(rawURL)
	if err != nil {
		return err
	}

	assets, pages := extractLinks(doc, baseURL)

	for _, asset := range assets {
		if !c.visited[asset] {
			c.visited[asset] = true

			resp, err := http.Get(asset)
			if err != nil {
				continue
			}

			data, err := io.ReadAll(resp.Body)
			resp.Body.Close()

			if err != nil {
				continue
			}

			_ = saveFile(asset, data)
		}
	}

	for _, page := range pages {
		u, err := url.Parse(page)
		if err != nil {
			continue
		}

		if u.Host != baseURL.Host {
			continue
		}

		_ = c.crawl(page, depth+1)
	}

	return nil
}

func saveFile(rawURL string, data []byte) error {
	u, err := url.Parse(rawURL)
	if err != nil {
		return err
	}

	path := filepath.Join("mirror", u.Host, u.Path)

	if u.Path == "" || u.Path == "/" {
		path = filepath.Join("mirror", u.Host, "index.html")
	}

	err = os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func extractLinks(doc *html.Node, base *url.URL) ([]string, []string) {
	var assets []string
	var pages []string

	var walk func(*html.Node)

	walk = func(n *html.Node) {
		if n.Type == html.ElementNode {

			switch n.Data {

			case "img":
				addAttr(n, "src", base, &assets)

			case "script":
				addAttr(n, "src", base, &assets)

			case "link":
				addAttr(n, "href", base, &assets)

			case "a":
				addAttr(n, "href", base, &pages)
			}
		}

		for child := n.FirstChild; child != nil; child = child.NextSibling {
			walk(child)
		}
	}

	walk(doc)

	return assets, pages
}

func addAttr(node *html.Node, key string, base *url.URL, target *[]string) {
	for _, attr := range node.Attr {

		if attr.Key != key {
			continue
		}

		ref, err := url.Parse(attr.Val)
		if err != nil {
			continue
		}

		*target = append(
			*target,
			base.ResolveReference(ref).String(),
		)
	}
}