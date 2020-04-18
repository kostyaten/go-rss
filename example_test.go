package rss

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ExampleParseFile() {
	rss, err := ParseFile("parser_test.rss")
	if err != nil {
		panic(err)
	}

	title := rss.Channel.Title
	description := rss.Channel.Description

	fmt.Println("Title:", title)
	fmt.Println("Description:", description)
}

func ExampleParseByte() {
	b, err := ioutil.ReadFile("parser_test.rss")
	if err != nil {
		panic(err)
	}

	rss, err := ParseByte(b)
	if err != nil {
		panic(err)
	}

	title := rss.Channel.Title
	description := rss.Channel.Description

	fmt.Println("Title:", title)
	fmt.Println("Description:", description)
}

func ExampleParseString() {
	rss, err := ParseString("<rss></rss>")
	if err != nil {
		panic(err)
	}

	title := rss.Channel.Title
	description := rss.Channel.Description

	fmt.Println("Title:", title)
	fmt.Println("Description:", description)
}

func ExampleParseURL() {
	rawurl := "http://static.userland.com/gems/backend/rssTwoExample2.xml"

	rss, err := ParseURL(rawurl, &http.Client{})
	if err != nil {
		panic(err)
	}

	title := rss.Channel.Title
	description := rss.Channel.Description

	fmt.Println("Title:", title)
	fmt.Println("Description:", description)
}
