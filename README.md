RSS 2.0 Parser
==============

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
