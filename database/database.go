package database

import "time"

type Player struct {
	ID         string
	Name       string
	Email      string
	InGameName string
}

type Match struct {
	ID       string
	Duration int
	MapName  string
	IsCustom bool
	BeginAt  time.Time
	EndAt    time.Time
}

type Cup struct {
	ID       string
	BeginAt  time.Time
	Winner   string // foreign key playerID
	GameMode string
}

type PlayerMatch struct {
	ID            string
	MatchID       string
	PlayerID      string
	DBNOs         int
	Assists       int
	DamageDealt   float64
	HeadshotKills int
	LongestKill   float64
	Kills         int
	Revives       int
	RideDistance  float64
	SwimDistance  float64
	WalkDistance  float64
	TimeSurvived  float64
	WinPlace      int
}

type CupMatch struct {
	ID      string
	CupID   string
	MatchID string
}

type Storage interface {
	GetPlayers() []Player
	CreatePlayer(Player) Player
	GetPlayer(Player) Player
	UpdatePlayer(Player) Player
	DeletePlayer(Player) Player

	GetMatches() []Match
	CreateMatch(Match) Match
	GetMatch(Match) Match
	UpdateMatch(Match) Match
	DeleteMatch(Match) Match

	GetCups() []Cup
	CreateCup(Cup) Cup
	GetCup(Cup) Cup
	UpdateCup(Cup) Cup
	DeleteCup(Cup) Cup

	GetPlayerMatches() []PlayerMatch
	CreatePlayerMatch(PlayerMatch) PlayerMatch
	GetPlayerMatch(PlayerMatch) PlayerMatch
	UpdatePlayerMatch(PlayerMatch) PlayerMatch
	DeletePlayerMatch(PlayerMatch) PlayerMatch

	GetCupMatches() []CupMatch
	CreateCupMatch(CupMatch) CupMatch
	GetCupMatch(CupMatch) CupMatch
	UpdateCupMatch(CupMatch) CupMatch
	DeleteCupMatch(CupMatch) CupMatch
}
