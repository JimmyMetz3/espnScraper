package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	teams := []string{}
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error: ", err.Error())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited: ", r.Request.URL)
	})

	c.OnHTML("a[data-track-nav_item]", func(e *colly.HTMLElement) {
		if e.Text == "" {
			return
		}
		teams = append(teams, e.Text)
		// fmt.Println(e.Text)
	})

	// c.Visit("https://www.espn.com/nfl/schedule/_/week/4/year/2024/seasontype/1")
	c.Visit("https://www.espn.com/nfl/schedule/_/week/1/year/2024/seasontype/2")
	total := 0
	for _, v := range teams[1:] {
		fmt.Println(v)
		total++
	}
	fmt.Println(total)
}
