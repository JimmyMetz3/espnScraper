package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type GameDay struct {
	AwayTeam Team
	HomeTeam Team
	Date     string
}

type Team struct {
	Wins   int
	Losses int
	Ties   int
}

var numOfGameDays int = 0
var days = []string{}

func addGameDay(index int, e *colly.HTMLElement) {
	if strings.Contains(e.ChildText("header.Card__Header"), "2024") {
		numOfGameDays++
	}
}

func getGameDayInfo(index int, e *colly.HTMLElement) {
	fmt.Println(e.ChildText("header.Card__Header"))
	fmt.Println(e.ChildText("div.ScoreCell__TeamName"))
	days = append(days, e.ChildText("header.Card__Header"))
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

	// Get Num of Days
	c.OnHTML("html", func(e *colly.HTMLElement) {
		e.ForEach("section.gameModules", addGameDay)
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
	// fmt.Println(teams)
	fmt.Println(numOfGameDays)
	fmt.Println(days)

	// Create new gamedays with days and teams

}
