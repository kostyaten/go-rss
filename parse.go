package rss

import (
	"compress/gzip"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Parse filename
func ParseFile(filename string) (*Rss, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return ParseByte(b)
}

// Parse bytes
func ParseByte(b []byte) (*Rss, error) {
	r := &Rss{}
	err := xml.Unmarshal(b, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Parse string
func ParseString(str string) (*Rss, error) {
	return ParseByte([]byte(str))
}

// Parse URL
func ParseURL(rawurl string, client *http.Client) (*Rss, error) {
	urlparse, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("GET", urlparse.String(), nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) "+
		"AppleWebKit/537.36 (KHTML, like Gecko) "+
		"Chrome/80.0.3987.163 Safari/537.36")

	request.Header.Set("Accept-Encoding", "gzip")
	request.Header.Set("Cache-Control", "max-age=0")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Check that the server actually sent compressed data
	var reader io.ReadCloser
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(response.Body)
		defer reader.Close()
	default:
		reader = response.Body
	}

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return ParseByte(b)
}

// Parse URL bulk
func ParseBulk(rawurls []string, client *http.Client, options *BulkOptions) []Rss {

	urlChan := make(chan string, options.buffer_chan)
	resChan := make(chan *Rss, options.buffer_chan)

	var sliceRss []Rss

	for i := 0; i < options.maxgoroutine; i++ {
		go func(urlChan chan string, resChan chan *Rss, id int) {
			for {
				rawurl, ok := <-urlChan
				if ok {
					feed, err := ParseURL(rawurl, client)
					if err != nil {
						panic(err)
					}
					resChan <- feed
				} else {
					break
				}
			}
		}(urlChan, resChan, i)
	}

	go func() {
		for _, v := range rawurls {
			urlChan <- v
		}
		close(urlChan)
	}()

	for i := 0; i < len(rawurls); i++ {
		res := <-resChan
		sliceRss = append(sliceRss, *res)
	}

	close(resChan)
	return sliceRss
}
