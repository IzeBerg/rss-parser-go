package rss

//////////////////////////////////////////////////////////////////////
//
// RSS structs that implements as per the RSS 2.0 specification.
// See http://www.rssboard.org/rss-specification
//
//////////////////////////////////////////////////////////////////////

import (
	"encoding/xml"
)

// <rss version="2.0">
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"channel,attr"`
	Channel Channel
}

// <channel>
//  <title>Liftoff News</title>
//  <link>http://liftoff.msfc.nasa.gov/</link>
//  <description>Liftoff to Space Exploration.</description>
//  <language>en-us</language>
//  <pubDate>Tue, 10 Jun 2003 04:00:00 GMT</pubDate>
//  <lastBuildDate>Tue, 10 Jun 2003 09:41:01 GMT</lastBuildDate>
//  <docs>http://blogs.law.harvard.edu/tech/rss</docs>
//  <generator>Weblog Editor 2.0</generator>
//  <managingEditor>editor@example.com</managingEditor>
//  <webMaster>webmaster@example.com</webMaster>
// </channel>
type Channel struct {
	XMLName        xml.Name `xml:"channel"`
	Title          string   `xml:"title"`
	Link           string   `xml:"link"`
	Description    string   `xml:"description"`
	Language       string   `xml:"language,omitempty"`
	Copyright      string   `xml:"copyright,omitempty"`
	ManagingEditor string   `xml:"managingEditor,omitempty"`
	WebMaster      string   `xml:"webMaster,omitempty"`
	PubDate        RSSDate  `xml:"pubDate,omitempty"`
	LastBuildDate  string   `xml:"lastBuildDate,omitempty"`
	Category       Category `xml:"category,omitempty"`
	Generator      string   `xml:"generator,omitempty"`
	Docs           string   `xml:"docs,omitempty"`
	Cloud          string   `xml:"cloud,omitempty"`
	TTL            string   `xml:"ttl,omitempty"`
	Image          string   `xml:"image,omitempty"`
	Rating         string   `xml:"rating,omitempty"`
	TextInput      string   `xml:"textInput,omitempty"`
	SkipHours      string   `xml:"skipHours,omitempty"`
	SkipDays       string   `xml:"skipDays,omitempty"`
	Items          []Item   `xml:"item"`
}

// <item>
//  <title>RSS Tutorial</title>
//  <link>http://www.w3schools.com/rss</link>
//  <description>New RSS tutorial on W3Schools</description>
// </item>
type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Author      string   `xml:"author,omitempty"`
	Category    Category `xml:"category,omitempty"`
	Comments    string   `xml:"comments,omitempty"`
	Enclosure   Enclosure
	Guid        Guid
	PubDate     RSSDate `xml:"pubDate"`
	Source      Source
}

// <source url="http://www.tomalak.org/links2.xml">Tomalak's Realm</source>
type Source struct {
	XMLName xml.Name `xml:"source"`
	Url     string   `xml:"url,attr"`
	Value   string   `xml:",innerxml"`
}

// <enclosure url="http://www.scripting.com/mp3s/weatherReportSuite.mp3"
//   length="12216320" type="audio/mpeg" />
type Enclosure struct {
	XMLName xml.Name `xml:"enclosure"`
	Url     string   `xml:"url,attr"`
	Length  int64    `xml:"length,attr"`
	Type    string   `xml:"type,attr"`
}

// <category>Grateful Dead</category>
// <category domain="http://www.fool.com/cusips">MSFT</category>
type Category struct {
	XMLName xml.Name `xml:"category"`
	Domain  string   `xml:"domain,attr"`
	Value   string   `xml:",innerxml"`
}

// <guid>http://some.server.com/weblogItem3207</guid>
// <guid isPermaLink="true">http://inessential.com/2002/09/01.php#a2</guid>
type Guid struct {
	XMLName   xml.Name `xml:"guid"`
	PermaLink string   `xml:"isPermaLink,attr"`
	Value     string   `xml:",innerxml"`
}

// <image>
//  <url>http://www.w3schools.com/images/logo.gif</url>
//  <title>W3Schools.com</title>
//  <link>http://www.w3schools.com</link>
// </image>
type Image struct {
	XMLName     xml.Name `xml:"image"`
	Url         string   `xml:"url"`
	Title       string   `xml:"title"`
	Link        string   `xml:"Link"`
	Width       int64    `xml:"width,omitempty"`
	Height      int64    `xml:"height,omitempty"`
	Description string   `xml:"description"`
}
