package rss

import (
	"fmt"
	"testing"
)

func TestFetchRss(t *testing.T) {
	//ParseRSS("http://www.rssboard.org/files/sample-rss-091.xml")
	//ParseRSS("http://www.rssboard.org/files/sample-rss-092.xml")
	rssObject, error := ParseRSS("http://www.rssboard.org/files/sample-rss-2.xml")

	if error != nil {
		fmt.Println(rssObject)
	}
}
