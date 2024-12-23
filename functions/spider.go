package functions

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/html"
)

var client = &http.Client{Timeout: 30 * time.Second}

// 遍历HTML节点，提取链接
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// 递归遍历子节点
		links = visit(links, c)
	}
	return links
}

func loadPage(url string) (*html.Node, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return html.Parse(resp.Body)
}

func FindLinks(url string) []string {
	doc, err := loadPage(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse html fail: %v\n", err)
		os.Exit(1)
	}
	links := visit(nil, doc)
	for _, link := range links {
		fmt.Println(link)
	}
	return links
}
