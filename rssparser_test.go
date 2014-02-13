package rss

import (
	"testing"
)

func TestFetchRss(t *testing.T) {
	ParseRSS("http://www.rssboard.org/files/sample-rss-091.xml")
	ParseRSS("http://www.rssboard.org/files/sample-rss-092.xml")
	ParseRSS("http://www.rssboard.org/files/sample-rss-2.xml")
}
