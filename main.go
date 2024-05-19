package main

import (
	"net/http"

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

func getGames(g *gin.Context) {
	g.IndentedJSON(http.StatusOK, games)
}

func main() {
	router := gin.Default()
	router.GET("/api/games/all", getGames)

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
