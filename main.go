package main

import (
  "fmt"
  "github.com/gocolly/colly"
)

func main() {
  fmt.Println("Starting...")
  c := colly.NewCollector(
    colly.AllowedDomains("www.fixkick.com"),
  )
  c.OnHTML("a[href]", func(e *colly.HTMLElement){
    link := e.Attr("href")
    fmt.Printf("Link found: %q -> %s\n", e.Text, link)
    c.Visit(e.Request.AbsoluteURL(link))
  })

  c.OnRequest(func(r *colly.Request){
    fmt.Println("Visiting", r.URL.String())
  })

  c.Visit("http://www.fixkick.com/")
}

