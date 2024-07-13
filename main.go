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

// Returns all games in JSON format
func getGames(client *gin.Context) {
	client.IndentedJSON(http.StatusOK, games)
}

// getGamesByID locates the game whose ID value matches the id
// parameter sent by the client, then returns that game as a response.
func getGamesByID(client *gin.Context) {
	id := client.Param("id")
	gameID, err := strconv.Atoi(id)
	if err != nil {
		client.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid game ID"})
		return
	}

	// loop over the list of games to find the id that matches the parameter
	for _, g := range games {
		if g.ID == gameID {
			client.IndentedJSON(http.StatusOK, g)
			return
		}
	}
	client.IndentedJSON(http.StatusNotFound, gin.H{"message": "game not found"})
}

// Get by games home team
func getGamesByHomeTeam(client *gin.Context) {
	name := client.Param("name")
	var teamGames []Game
	for _, g := range games {
		if g.HomeTeam.Name == name {
			teamGames = append(teamGames, g)
		}
	}
	if len(teamGames) > 0 {
		client.IndentedJSON(http.StatusOK, teamGames)
		return
	} else {
		client.IndentedJSON(http.StatusNotFound, gin.H{"message": "team not found"})
	}
}

// Get by games by away team
func getGamesByAwayTeam(client *gin.Context) {
	name := client.Param("name")
	var teamGames []Game
	for _, g := range games {
		if g.AwayTeam.Name == name {
			teamGames = append(teamGames, g)
		}
	}
	if len(teamGames) > 0 {
		client.IndentedJSON(http.StatusOK, teamGames)
		return
	} else {
		client.IndentedJSON(http.StatusNotFound, gin.H{"message": "team not found"})
	}
}

// Returns games by year
func getGamesByYear(client *gin.Context) {
	yearParam := client.Param("year")

	if len(yearParam) != 4 {
		client.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid year format"})
		return
	}

	year, err := strconv.Atoi(yearParam)

	currentYear := time.Now().Year()
	if err != nil || year < 1889 || year > currentYear {
		client.IndentedJSON(http.StatusBadRequest, gin.H{"message": "year out of range"})
		return
	}

	var yearGames []Game
	for _, g := range games {
		gameDate, err := time.Parse("2006-01-02", g.Date)
		if err != nil {
			continue
		}
		if gameDate.Year() == year {
			yearGames = append(yearGames, g)

		}
	}

	if len(yearGames) > 0 {
		client.IndentedJSON(http.StatusOK, yearGames)
		return
	} else {
		client.IndentedJSON(http.StatusNotFound, gin.H{"message": "no games found for that year"})
	}
}

func main() {
	router := gin.Default()

	// Load all templates from the templates directory
	router.SetHTMLTemplate(template.Must(template.ParseGlob("templates/*.html")))

	// Serve static files from the static directory
	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
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

		c.HTML(http.StatusOK, "index.html", gin.H{
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

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
