package main

import (
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/seungwooj/learngo/scrapper"
)

const fileName string= "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove("jobs.csv")
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	location := strings.ToLower(scrapper.CleanString(c.FormValue("location")))
	scrapper.Scrape(term, location)
	return c.Attachment("jobs.csv", "jobs.csv")
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	// scrapper.Scrape("term", "location")
	e.Logger.Fatal(e.Start(":1323"))
}