package rss

//////////////////////////////////////////////////////////////////////
//
// RSSDate interface primarily provides the tools to parse and format
// RFC1123 date strings.
//
// Example: "Tue, 10 Jun 2003 04:00:00 GMT"
//
//////////////////////////////////////////////////////////////////////

import (
	"time"
)

type RSSDate string

// Parse a given RFC1123 formatted string and return a time object.
func (self RSSDate) Parse() (time.Time, error) {
	return time.Parse(time.RFC1123, string(self))
}
