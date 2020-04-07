# govid19 - Worldometer COVID-19 scrapper

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gabrielmbmb/govid19)

**govid19** is a very simple scraper of [Worldometer COVID-19](https://www.worldometers.info/coronavirus/) developed with Go. 
This is my very first Go package. Since the start of the isolation, I've been quite bored so I decided to learn Go... *sigh*
As you will see, the package is not very big because its main purpose was to learn how to structure a Go package

## Usage

```go
package main

import "github.com/gabrielmbmb/govid19"

func main() {
	countries := govid19.Scrape()
	err := govid19.WriteToCSV(countries)
	if err != nil {
		panic(err)
	}
}
```
