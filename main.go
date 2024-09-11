package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

type Week struct {
	Week     int
	Gamedays DayOfWeek
}

type DayOfWeek struct {
	Day   time.Weekday
	Games Matchup
}

type Matchup struct {
	AwayTeam string
	HomeTeam string
}

// var numOfGameDays int = 0

// var days = []string{}

// func addGameDay(index int, e *colly.HTMLElement) {
// 	if strings.Contains(e.ChildText("header.Card__Header"), "2024") {
// 		// numOfGameDays++
// 	}
// }

func getGameDayInfo(index int, e *colly.HTMLElement) {
	fmt.Println(e.ChildText("header.Card__Header"))
	fmt.Println(e.ChildText("div.ScoreCell__TeamName"))
	fmt.Println(e.ChildText("div>span.ScoreboardCell__Record"))
	// days = append(days, e.ChildText("header.Card__Header"))
}

func main() {
	c := colly.NewCollector()

	// teams := []string{}

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Error: ", err.Error())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited: ", r.Request.URL)
	})

	// Get Num of Days
	// c.OnHTML("html", func(e *colly.HTMLElement) {
	// 	e.ForEach("section.gameModules", addGameDay)
	// })

	// Get Date with ForEach
	c.OnHTML("html", func(e *colly.HTMLElement) {
		e.ForEach("section.gameModules", getGameDayInfo)
	})

	// Get Teams names
	// c.OnHTML("div.ScoreCell__TeamName", func(e *colly.HTMLElement) {
	// 	teams = append(teams, e.Text)
	// })

	c.Visit("https://www.espn.com/nfl/scoreboard/_/week/1/year/2024/seasontype/2")

	// fmt.Println(teams)
	// fmt.Println(numOfGameDays)
	// fmt.Println(days)

	// Create new gamedays with days and teams

}
