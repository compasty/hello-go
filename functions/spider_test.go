package functions

import "testing"

func TestFindLinks(t *testing.T) {
	// Test the findLinks function.
	links := FindLinks("https://www.baidu.com")
	if len(links) == 0 {
		t.Errorf("findLinks() returned no links")
	}
}
