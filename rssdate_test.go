package rss

import (
	"testing"
)

func TestRSSDate(t *testing.T) {

	var d RSSDate = "Tue, 10 Jun 2002 22:11:22 EDT"
	time, _ := d.Parse()
	if time.Year() != 2002 || time.Month() != 6 || time.Day() != 10 ||
		time.Hour() != 22 || time.Minute() != 11 || time.Second() != 22 {
		t.Fail()
	}

	name, _ := time.Zone()

	if name != "EDT" {
		t.Fail()
	}
}
