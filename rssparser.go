package rss

//////////////////////////////////////////////////////////////////////
//
// RSSParser provides functions to parse RSS feed via HTTP URLs.
//
//////////////////////////////////////////////////////////////////////

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func ParseRSS(url string) (*RSS, error) {

	response, err1 := http.Get(url)
	if err1 != nil {
		log.Println("Unable to retrieve via HTTP", url)
		return nil, err1
	}
	defer response.Body.Close()

	data, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		log.Println("Unable to read bytes from socket")
		return nil, err2
	}

	rss := RSS{}
	err3 := xml.Unmarshal(data, &rss)
	if err3 != nil {
		return nil, err3
	}
	return &rss, os.ErrNotExist
}
