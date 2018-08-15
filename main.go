package main

import (
  "fmt"
  "github.com/gocolly/colly"
  "encoding/csv"
  "os"
  "log"
)

func main() {
  titles := []string{"Text", "Href"}
  record := []string{"",""}
  outfile, err := os.Create("./resultsfile.csv")
  if err != nil {
    log.Fatal(err)
  }
  w := csv.NewWriter(outfile)
  w.Write(titles)

  fmt.Println("Starting...")
  c := colly.NewCollector(
    colly.AllowedDomains("www.fixkick.com"),
  )

  c.OnHTML("a[href]", func(e *colly.HTMLElement){
    link := e.Attr("href")
    fmt.Printf("Link found: %q -> %s\n", e.Text, link)
    c.Visit(e.Request.AbsoluteURL(link))
    record[0] = e.Text
    record[1] = link
    w.Write(record)
    w.Flush()
  })

  c.OnRequest(func(r *colly.Request){
    fmt.Println("Visiting", r.URL.String())
  })

  c.Visit("http://www.fixkick.com/")
}

