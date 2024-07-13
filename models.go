package main

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
