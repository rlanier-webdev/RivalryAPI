package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Team struct {
	Name string `json:"name"`
}

type Game struct {
	ID       int    `json:"id"`
	HomeTeam *Team  `json:"homeTeam"`
	AwayTeam *Team  `json:"awayTeam"`
	Date     string `json:"date"`
	Score    *Score `json:"score"`
	Notes    string `json:"notes"`
}

type Score struct {
	HomeTeamScore int `json:"homeTeamScore"`
	AwayTeamScore int `json:"awayTeamScore"`
}

var games = []Game{
	{ID: 1, HomeTeam: &Team{Name: "City"}, AwayTeam: &Team{Name: "Poly"}, Date: "2000-11-01", Score: &Score{HomeTeamScore: 10, AwayTeamScore: 7}, Notes: ""},
	{ID: 2, HomeTeam: &Team{Name: "Poly"}, AwayTeam: &Team{Name: "City"}, Date: "2001-11-01", Score: &Score{HomeTeamScore: 20, AwayTeamScore: 10}, Notes: ""},
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

// keep for now, but figure out why this would be used
func getGamesByTeam(client *gin.Context) {
	name := client.Param("name")
	var teamGames []Game
	for _, g := range games {
		if g.HomeTeam.Name == name || g.AwayTeam.Name == name {
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
	router.GET("/api/games/all", getGames)
	router.GET("/api/games/:id", getGamesByID)
	router.GET("/api/games/team/:name", getGamesByTeam)
	router.GET("/api/games/year/:year", getGamesByYear)

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
