package rss

import (
	"encoding/xml"
	"net/url"
	"time"
)

//https://validator.w3.org/feed/docs/rss2.html

type Rss struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title          string    `xml:"title"`       // required
	Link           *URL      `xml:"link"`        // required
	Description    string    `xml:"description"` // required
	Language       string    `xml:"language,omitempty"`
	Copyright      string    `xml:"copyright,omitempty"`
	ManagingEditor string    `xml:"managingEditor,omitempty"`
	WebMaster      string    `xml:"webMaster,omitempty"`
	PubDate        *DateTime `xml:"pubDate,omitempty"`
	LastBuildDate  *DateTime `xml:"lastBuildDate,omitempty"`
	Category       []string  `xml:"category,omitempty"`
	Generator      string    `xml:"generator,omitempty"`
	Docs           string    `xml:"docs,omitempty"`
	TTL            int       `xml:"ttl,omitempty"`
	Image          *Image    `xml:"image,omitempty"`
	Item           []*Item   `xml:"item"`
}

type Item struct {
	Title       string       `xml:"title,omitempty"`
	Description string       `xml:"description,omitempty"`
	Link        *URL         `xml:"link,omitempty"`
	Author      string       `xml:"author,omitempty"`
	Category    []string     `xml:"category,omitempty"`
	Comments    string       `xml:"comments,omitempty"`
	Enclosure   []*Enclosure `xml:"enclosure,omitempty"`
	Guid        string       `xml:"guid,omitempty"`
	PubDate     string       `xml:"pubDate,omitempty"`
	Source      string       `xml:"source,omitempty"`
}

type Enclosure struct {
	URL    string `xml:"url,omitempty,attr"`
	Length int    `xml:"length,omitempty,attr"`
	Type   string `xml:"type,omitempty,attr"`
}

type Image struct {
	URL    *URL   `xml:"url,omitempty"`
	Title  string `xml:"title,omitempty"`
	Link   *URL   `xml:"link,omitempty"`
	Width  int    `xml:"width,omitempty"`
	Height int    `xml:"height,omitempty"`
}

type DateTime struct {
	Time time.Time
}

type URL struct {
	URL *url.URL
}

func (u *URL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}

	res, err := url.Parse(v)
	if err != nil {
		return err
	}

	u.URL = res
	return nil
}

func (t *DateTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}

	parse, err := time.Parse(time.RFC1123, v)
	if err != nil {
		return err
	}

	t.Time = parse
	return nil
}
