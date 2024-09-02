package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

type GameDay struct {
	MatchUp struct{}
	date    time.Time
}

type MatchUp struct {
	AwayTeam string
	HomeTeam string
}

func main() {
	c := colly.NewCollector()

	teams := []string{}
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error: ", err.Error())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited: ", r.Request.URL)
	})

	// Get Date with ForEach
	c.OnHTML("html", func(e *colly.HTMLElement) {
		e.ForEach("section.gameModules", getGameDayInfo)
	})

	// Get Teams names
	c.OnHTML("div.ScoreCell__TeamName", func(e *colly.HTMLElement) {
		teams = append(teams, e.Text)
	})

	c.Visit("https://www.espn.com/nfl/scoreboard/_/week/1/year/2024/seasontype/2")
	fmt.Println(teams)
}

func getGameDayInfo(index int, h *colly.HTMLElement) {
	fmt.Println(h.ChildText("header.Card__Header"))
	fmt.Println(h.ChildText("div.ScoreCell__TeamName"))
}
