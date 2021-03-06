rss-parser-go
=============

RSS 2.0 parser based in GO.  Includes convenience function to fetch RSS via HTTP URL.  For the entire RSS specification, see http://www.rssboard.org/rss-specification

See the domain mode of the RSS object: https://github.com/RealJK/rss-parser-go/blob/master/rss.go

Installation
============

    > go get github.com/RealJK/rss-parser-go

Using the Library
=================

    package main

    import (
	    "fmt"
	    "github.com/RealJK/rss-parser-go"
    )

    func main() {

	    rssObject, err := rss.ParseRSS("http://www.rssboard.org/files/sample-rss-2.xml")
	    if err != nil {

		    fmt.Printf("Title           : %s\n", rssObject.Channel.Title)
    		fmt.Printf("Generator       : %s\n", rssObject.Channel.Generator)
    		fmt.Printf("PubDate         : %s\n", rssObject.Channel.PubDate)
    		fmt.Printf("LastBuildDate   : %s\n", rssObject.Channel.LastBuildDate)
    		fmt.Printf("Description     : %s\n", rssObject.Channel.Description)

    		fmt.Printf("Number of Items : %d\n", len(rssObject.Channel.Items))

	    	for v := range rssObject.Channel.Items {
	    		item := rssObject.Channel.Items[v]
    			fmt.Println()
	    		fmt.Printf("Item Number : %d\n", v)
	    		fmt.Printf("Title       : %s\n", item.Title)
	    		fmt.Printf("Link        : %s\n", item.Link)
	    		fmt.Printf("Description : %s\n", item.Description)
	    		fmt.Printf("Guid        : %s\n", item.Guid.Value)
    		}
	    }
    }
