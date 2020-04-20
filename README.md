RSS 2.0 Parser
==============

[![Go](https://github.com/kostya-ten/go-rss/workflows/Go/badge.svg?branch=master)](https://github.com/kostya-ten/go-rss/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/kostya-ten/go-rss)](https://goreportcard.com/report/github.com/kostya-ten/go-rss)
[![GoDoc](https://godoc.org/github.com/kostya-ten/go-rss?status.svg&style=flat)](https://pkg.go.dev/github.com/kostya-ten/go-rss)

## Requirements

    Golang 1.13+, 1.14+

## Getting It

You can get go-rss by using

    $ go get -u github.com/kostya-ten/go-rss


### Parse rss file
```go

  import "github.com/kostya-ten/go-rss"

  func main() {
    feed, err := rss.ParseFile(filename)
    if err != nil {
      panic(err)
    }

    title := feed.Channel.Title
    fmt.Println(title)
  }
```


### Parse rss url 
```go

  import "github.com/kostya-ten/go-rss"

  func main() {
    feed, err := rss.ParseURL("http://static.userland.com/gems/backend/rssTwoExample2.xml", &http.Client{})
    if err != nil {
      panic(err)
    }

    title := feed.Channel.Title
    fmt.Println(title)
  }
```

### Bulk parse rss url
```go
  import "github.com/kostya-ten/go-rss"

  func main() {
    urls := []string{
      "https://lenta.ru/rss",
      "https://www.interfax.ru/rss.asp",
      "https://ria.ru/export/rss2/index.xml",
      "http://static.feed.rbc.ru/rbc/logical/footer/news.rss",
      "http://tass.ru/rss/v2.xml",
      "https://www.vesti.ru/vesti.rss",
    }

    resultRss := rss.ParseBulk(urls, &http.Client{}, &BulkOptions{maxgoroutine: 10, buffer_chan: 10})
    for _, v := range resultRss {
      fmt.Println(v.Channel.Title)
    }
  }
```
