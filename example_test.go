package rss

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ExampleParseFile() {
	feed, err := ParseFile("parser_test.rss")
	if err != nil {
		panic(err)
	}

	title := feed.Channel.Title
	description := feed.Channel.Description

	fmt.Println("Title:", title)
	fmt.Println("Description:", description)
}

func ExampleParseByte() {
	b, err := ioutil.ReadFile("parser_test.rss")
	if err != nil {
		panic(err)
	}

	feed, err := ParseByte(b)
	if err != nil {
		panic(err)
	}

	title := feed.Channel.Title
	description := feed.Channel.Description

	fmt.Println("Title:", title)
	fmt.Println("Description:", description)
}

func ExampleParseString() {
	feed, err := ParseString("<rss></rss>")
	if err != nil {
		panic(err)
	}

	title := feed.Channel.Title
	description := feed.Channel.Description

	fmt.Println("Title:", title)
	fmt.Println("Description:", description)
}

func ExampleParseURL() {
	rawurl := "http://static.userland.com/gems/backend/rssTwoExample2.xml"

	feed, err := ParseURL(rawurl, &http.Client{})
	if err != nil {
		panic(err)
	}

	title := feed.Channel.Title
	description := feed.Channel.Description

	fmt.Println("Title:", title)
	fmt.Println("Description:", description)
}
