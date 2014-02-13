package rss

//////////////////////////////////////////////////////////////////////
//
// RSSParser provides functions to parse RSS feed via HTTP URLs.
//
//////////////////////////////////////////////////////////////////////

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

func ParseRSS(url string) (*RSS, error) {

	rss := RSS{}
	response, err1 := http.Get(url)
	if err1 != nil {
		return nil, err1
	}
	defer response.Body.Close()

	data, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		return nil, err2
	}

	err3 := xml.Unmarshal(data, &rss)
	if err3 != nil {
		return nil, err3
	}
	return &rss, nil
}
