package rss

import (
	"encoding/xml"
	"io/ioutil"
	"testing"
)

func TestRSS091(t *testing.T) {
	rss := RSS{}
	data, _ := ioutil.ReadFile("tests/sample-rss-091.xml")
	xml.Unmarshal(data, &rss)
}

func TestRSS092(t *testing.T) {
	rss := RSS{}
	data, _ := ioutil.ReadFile("tests/sample-rss-092.xml")
	xml.Unmarshal(data, &rss)
}

func TestRSS200(t *testing.T) {
	rss := RSS{}
	data, _ := ioutil.ReadFile("tests/sample-rss-2.xml")
	xml.Unmarshal(data, &rss)
}

//<channel>
//	<title>title</title>
//	<link>link</link>
//	<description>description</description>
//	<language>language</language>
//	<pubDate>Tue, 10 Jun 2003 04:00:00 GMT</pubDate>
//	<generator>generator</generator>
//	<managingEditor>managingEditor</managingEditor>
//	<webMaster>webmaster@example.com</webMaster>
//	<item>
//		<title>item title</title>
//		<link>http://liftoff.msfc.nasa.gov/news/2003/news-starcity.asp</link>
//		<description>item description</description>
//		<pubDate>Tue, 03 Jun 2003 09:39:21 GMT</pubDate>
//		<guid>item guid</guid>
//	</item>
//</channel>
func TestStubbed(t *testing.T) {
	rss := RSS{}
	data, _ := ioutil.ReadFile("tests/stubbed-rss-2.xml")
	xml.Unmarshal(data, &rss)

	if rss.Channel.Title != "title" ||
		rss.Channel.Link != "link" ||
		rss.Channel.Description != "description" ||
		rss.Channel.Language != "language" ||
		rss.Channel.Generator != "generator" {
		t.Error("Wrong data")
	}

	time, _ := rss.Channel.PubDate.Parse()

	if time.Month() != 6 ||
		time.Year() != 2003 ||
		time.Hour() != 4 ||
		time.Minute() != 0 {
		t.Error("Wrong data")
	}

	if len(rss.Channel.Items) != 1 {
		t.Error("Wrong data")
	}

	if rss.Channel.Items[0].Title != "item title" ||
		rss.Channel.Items[0].Description != "item description" ||
		rss.Channel.Items[0].Guid.Value != "item guid" {
		t.Error("Wrong data")
	}
}
