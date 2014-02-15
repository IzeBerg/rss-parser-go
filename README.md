rss-parser-go
=============

RSS 2.0 parser based in GO.  Includes convenience function to fetch RSS via HTTP URL.  For the entire RSS specification, see http://www.rssboard.org/rss-specification

Usage
=====

Step 1: Download the library

    > go get github.com/RealJK/rss-parser-go
    
Step 2: Use the library

    package main

    import(
        "fmt"
        "github.com/RealJK/rss-parser-go"
    )

    func main() {

        rssObject, err := rss.ParseRSS("http://www.rssboard.org/files/sample-rss-2.xml")
        if err != nil {
                fmt.Println(rssObject)
        }
    }
