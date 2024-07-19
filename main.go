package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
)

var games []Game

func init() {
	// Read data from JSON file
	data, err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	// Unmarshal JSON data into games slice
	err = json.Unmarshal(data, &games)
	if err != nil {
		panic(err)
	}
}

func main() {
	router := gin.Default()

	// Load all templates from the templates directory
	router.LoadHTMLGlob("templates/*.html")

	// Serve static files from the static directory
	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		var message string

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title":   "Welcome",
			"Message": message,
		})
	})

	router.GET("/all", func(c *gin.Context) {
		results := games
		var message string

		c.HTML(http.StatusOK, "all-games.html", gin.H{
			"Title":   "All Games",
			"Results": results,
			"Message": message,
		})
	})

	router.GET("/search", func(c *gin.Context) {
		searchType := c.Query("searchType")
		query := c.Query("query")
		var results []Game
		var message string

		switch searchType {
		case "id":
			id, err := strconv.Atoi(query)
			if err == nil {
				for _, g := range games {
					if g.ID == id {
						results = append(results, g)
					}
				}
			}
		case "home":
			for _, g := range games {
				if g.HomeTeam.Name == query {
					results = append(results, g)
				}
			}
		case "away":
			for _, g := range games {
				if g.AwayTeam.Name == query {
					results = append(results, g)
				}
			}
		case "year":
			year, err := strconv.Atoi(query)
			if err == nil {
				for _, g := range games {
					gameDate, err := time.Parse("2006-01-02", g.Date)
					if err == nil && gameDate.Year() == year {
						results = append(results, g)
					}
				}
			}
		}

		if len(results) == 0 {
			message = "No results found"
		}

		c.HTML(http.StatusOK, "search.html", gin.H{
			"Title":   "Search",
			"Results": results,
			"Message": message,
		})
	})

	router.GET("/api/games/all", getGames)
	router.GET("/api/games/:id", getGamesByID)
	router.GET("/api/games/home/:name", getGamesByHomeTeam)
	router.GET("/api/games/away/:name", getGamesByAwayTeam)
	router.GET("/api/games/year/:year", getGamesByYear)

	// Serve the documentation.html file
	router.GET("/docs", func(c *gin.Context) {
		mdContent, err := os.ReadFile("README.md")
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to read documentation")
			return
		}
		htmlContent := blackfriday.MarkdownCommon(mdContent)

		c.HTML(http.StatusOK, "documentation.html", gin.H{
			"Content": template.HTML(htmlContent),
		})
	})

	err := router.Run("localhost:1889")
	if err != nil {
		panic(err)
	}
}
