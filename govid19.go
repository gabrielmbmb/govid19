package govid19

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/gocolly/colly"
	"os"
	"strconv"
	"strings"
	"time"
)

type Country struct {
	Name            string `csv:"country_other"`
	TotalCases      int    `csv:"total_cases"`
	NewCases        int    `csv:"new_cases"`
	TotalDeaths     int    `csv:"total_deaths"`
	NewDeaths       int    `csv:"new_deaths"`
	TotalRecovered  int    `csv:"total_recovered"`
	ActiveCases     int    `csv:"active_cases"`
	SeriousCritical int    `csv:"seriousor_critical"`
	TotalCases1MPop int    `csv:"total_cases_1m_pop"`
	Deaths1MPop     int    `csv:"deaths_1m_pop"`
	TotalTests      int    `csv:"total_tests"`
	Tests1MPop      int    `csv:"tests_1m_pop"`
}

func parseToInt(str string) int {
	if val, err := strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(str, ",", ""), " ", ""), 10, 64); err == nil {
		return int(val)
	}
	return 0
}

func WriteToCSV(countries []*Country) error {
	fn := fmt.Sprintf("covid-%s.csv", time.Now().Format(time.RFC3339))
	file, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Printf("Saving scrapped data in %s...\n", fn)
	err = gocsv.MarshalFile(&countries, file)
	if err != nil {
		return err
	}

	return nil
}

func Scrape() []*Country {
	url := "https://www.worldometers.info/coronavirus/"
	countries := []*Country{}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("#main_table_countries_today > tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, e *colly.HTMLElement) {
			if e.Attr("class") != "total_row_world" {
				country := Country{}
				e.ForEach("td", func(col int, e *colly.HTMLElement) {
					switch col {
					case 0:
						if val, err := e.DOM.Children().Html(); err == nil {
							country.Name = val
						}
					case 1:
						if val, err := e.DOM.Html(); err == nil {
							country.TotalCases = parseToInt(val)
						}
					case 2:
						if val, err := e.DOM.Html(); err == nil {
							country.NewCases = parseToInt(val)
						}
					case 3:
						if val, err := e.DOM.Html(); err == nil {
							country.TotalDeaths = parseToInt(val)
						}
					case 4:
						if val, err := e.DOM.Html(); err == nil {
							country.NewDeaths = parseToInt(val)
						}
					case 5:
						if val, err := e.DOM.Html(); err == nil {
							country.TotalRecovered = parseToInt(val)
						}
					case 6:
						if val, err := e.DOM.Html(); err == nil {
							country.ActiveCases = parseToInt(val)
						}
					case 7:
						if val, err := e.DOM.Html(); err == nil {
							country.SeriousCritical = parseToInt(val)
						}
					case 8:
						if val, err := e.DOM.Html(); err == nil {
							country.TotalCases1MPop = parseToInt(val)
						}
					case 9:
						if val, err := e.DOM.Html(); err == nil {
							country.Deaths1MPop = parseToInt(val)
						}
					case 10:
						if val, err := e.DOM.Html(); err == nil {
							country.TotalTests = parseToInt(val)
						}
					case 11:
						if val, err := e.DOM.Html(); err == nil {
							country.Tests1MPop = parseToInt(val)
						}
					case 12:
						if val, err := e.DOM.Html(); err == nil {
							country.Tests1MPop = parseToInt(val)
						}
					default:
						fmt.Printf("Found extranous column in position %d\n", col)
					}
				})
				countries = append(countries, &country)
			}
		})
	})

	c.Visit(url)
	return countries
}
