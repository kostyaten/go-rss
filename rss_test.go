package rss

import (
	"io/ioutil"
	"net/http"
	"testing"
)

const filename = "test.rss"

func TestParseFile(t *testing.T) {
	rss, err := ParseFile(filename)
	if err != nil {
		t.Error(err)
	}

	title := rss.Channel.Title
	description := rss.Channel.Description

	if title != "Lorem ipsum feed for an interval of 1 minutes" {
		t.Error("Invalid error parse title")
	}

	if description != "This is a constantly updating lorem ipsum feed" {
		t.Error("Invalid error parse description")
	}
}

func TestParseByte(t *testing.T) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error(err)
	}

	rss, err := ParseByte(b)
	if err != nil {
		t.Error(err)
	}

	title := rss.Channel.Title
	description := rss.Channel.Description

	if title != "Lorem ipsum feed for an interval of 1 minutes" {
		t.Error("Invalid error parse title")
	}

	if description != "This is a constantly updating lorem ipsum feed" {
		t.Error("Invalid error parse description")
	}
}

func TestParseString(t *testing.T) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Error(err)
	}

	rss, err := ParseString(string(b))
	if err != nil {
		t.Error(err)
	}

	title := rss.Channel.Title
	description := rss.Channel.Description

	if title != "Lorem ipsum feed for an interval of 1 minutes" {
		t.Error("Invalid error parse title")
	}

	if description != "This is a constantly updating lorem ipsum feed" {
		t.Error("Invalid error parse description")
	}
}

func TestParseURL(t *testing.T) {
	rss, err := ParseURL("http://static.userland.com/gems/backend/rssTwoExample2.xml", &http.Client{})
	if err != nil {
		t.Error(err)
	}

	title := rss.Channel.Title
	description := rss.Channel.Description

	if title != "Scripting News" {
		t.Error("Invalid error parse title")
	}

	if description != "A weblog about scripting and stuff like that." {
		t.Error("Invalid error parse description")
	}
}

func TestChannelImage(t *testing.T) {
	rss, err := ParseFile(filename)
	if err != nil {
		t.Error(err)
	}

	if rss.Channel.Image.URL.URL.String() != "https://lenta.ru/images/small_logo.png" {
		t.Error("Invalid error parse Channel->Image->URL")
	}

	if rss.Channel.Image.Title != "Lenta.ru" {
		t.Error("Invalid error parse Channel->Image->Title")
	}

	if rss.Channel.Image.Link.URL.String() != "https://lenta.ru" {
		t.Error("Invalid error parse Channel->Image->Link")
	}

	if rss.Channel.Image.Height != 22 {
		t.Error("Invalid error parse Channel->Image->Height")
	}

	if rss.Channel.Image.Width != 134 {
		t.Error("Invalid error parse Channel->Image->Width")
	}

}
