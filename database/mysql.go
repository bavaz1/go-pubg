package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/xid"
)

type MySQL struct {
	db *sql.DB
}

func New() Storage {
	db, err := sql.Open("mysql", "root:toor@/pubg?parseTime=true")
	checkErr(err)

	m := new(MySQL)
	m.db = db
	return m
}

func (m MySQL) GetPlayers() []Player {
	var (
		player  Player
		players []Player
	)

	rows, err := m.db.Query("select * from Players;")

	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&player.ID, &player.Name, &player.Email, &player.InGameName)
		players = append(players, player)
		checkErr(err)

	}

	defer rows.Close()
	return players
}

func (m MySQL) CreatePlayer(player Player) Player {
	player.ID = xid.New().String()
	stmt, err := m.db.Prepare("insert into Players(id, name, email, inGameName) VALUES (?, ?, ?, ?)")
	checkErr(err)

	_, err = stmt.Exec(player.ID, player.Name, player.Email, player.InGameName)
	checkErr(err)

	return player
}

func (m MySQL) GetPlayer(player Player) Player {
	rows, err := m.db.Query("select * from Players where id = ?;", player.ID)

	checkErr(err)

	var p Player
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Name, &p.Email, &p.InGameName)
		checkErr(err)

	}

	defer rows.Close()
	return p
}

func (m MySQL) UpdatePlayer(player Player) Player {
	stmt, err := m.db.Prepare("update Players set name=?, email=?, inGameName=? where id=?")
	checkErr(err)

	_, err = stmt.Exec(player.Name, player.Email, player.InGameName, player.ID)
	checkErr(err)

	return player
}

func (m MySQL) DeletePlayer(player Player) Player {
	stmt, err := m.db.Prepare("delete from Players where id=?")
	checkErr(err)

	_, err = stmt.Exec(player.ID)
	checkErr(err)

	return player
}

func (m MySQL) GetMatches() []Match {
	var (
		match   Match
		matches []Match
	)

	rows, err := m.db.Query("select * from `Match`;")

	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&match.ID, &match.Duration, &match.MapName, &match.IsCustom, &match.BeginAt)
		matches = append(matches, match)
		checkErr(err)

	}

	defer rows.Close()
	return matches
}

func (m MySQL) CreateMatch(match Match) Match {
	match.ID = xid.New().String()
	stmt, err := m.db.Prepare("insert into `Match`(id, duration, mapName, isCustom, beginAt) VALUES (?, ?, ?, ?, ?)")
	checkErr(err)

	_, err = stmt.Exec(match.ID, match.Duration, match.MapName, match.IsCustom, match.BeginAt)
	checkErr(err)

	return match
}

func (m MySQL) GetMatch(match Match) Match {
	rows, err := m.db.Query("select * from `Match` where id = ?;", match.ID)

	checkErr(err)

	var matchRes Match
	for rows.Next() {
		err = rows.Scan(&matchRes.ID, &matchRes.Duration, &matchRes.MapName, &matchRes.IsCustom, &matchRes.BeginAt)
		checkErr(err)

	}

	defer rows.Close()
	return matchRes
}

func (m MySQL) UpdateMatch(match Match) Match {
	stmt, err := m.db.Prepare("update `Match` set duration=?, mapName=?, isCustom=?, beginAt=? where id=?")
	checkErr(err)

	_, err = stmt.Exec(match.Duration, match.MapName, match.IsCustom, match.BeginAt, match.ID)
	checkErr(err)

	return match
}

func (m MySQL) DeleteMatch(match Match) Match {
	stmt, err := m.db.Prepare("delete from `Match` where id=?")
	checkErr(err)

	_, err = stmt.Exec(match.ID)
	checkErr(err)

	return match
}

func (m MySQL) GetCups() []Cup {
	var (
		cup  Cup
		cups []Cup
	)
	var winner sql.NullString

	rows, err := m.db.Query("select * from Cups;")

	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&cup.ID, &cup.BeginAt, &winner, &cup.GameMode)
		if winner.Valid {
			cup.Winner = winner.String
		}
		cups = append(cups, cup)
		checkErr(err)

	}

	defer rows.Close()
	return cups
}

func (m MySQL) CreateCup(cup Cup) Cup {
	cup.ID = xid.New().String()
	stmt, err := m.db.Prepare("insert into Cups(id, beginAt, winner, gameMode) VALUES (?, ?, ?, ?)")
	checkErr(err)

	_, err = stmt.Exec(cup.ID, cup.BeginAt, nil, cup.GameMode)
	checkErr(err)

	return cup
}

func (m MySQL) GetCup(cup Cup) Cup {
	rows, err := m.db.Query("select * from Cups where id = ?;", cup.ID)

	checkErr(err)

	var c Cup
	var winner sql.NullString
	for rows.Next() {
		err = rows.Scan(&c.ID, &c.BeginAt, &winner, &c.GameMode)
		checkErr(err)

	}
	if winner.Valid {
		c.Winner = winner.String
	}

	defer rows.Close()
	return c
}

func (m MySQL) UpdateCup(cup Cup) Cup {
	stmt, err := m.db.Prepare("update Cups set beginAt=?, winner=?, gameMode=? where id=?")
	checkErr(err)

	_, err = stmt.Exec(cup.BeginAt, cup.Winner, cup.GameMode, cup.ID)
	checkErr(err)

	return cup
}

func (m MySQL) DeleteCup(cup Cup) Cup {
	stmt, err := m.db.Prepare("delete from Cups where id=?")
	checkErr(err)

	_, err = stmt.Exec(cup.ID)
	checkErr(err)

	return cup
}

func (m MySQL) GetPlayerMatches(playerMatch PlayerMatch) []PlayerMatch {
	var (
		playerMatches []PlayerMatch
	)

	rows, err := m.db.Query("select * from PlayerMatch where playerID = ?;", playerMatch.PlayerID)

	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&playerMatch.ID, &playerMatch.MatchID, &playerMatch.PlayerID, &playerMatch.DBNOs, &playerMatch.Assists, &playerMatch.DamageDealt, &playerMatch.HeadshotKills, &playerMatch.LongestKill, &playerMatch.Kills, &playerMatch.Revives, &playerMatch.RideDistance, &playerMatch.SwimDistance, &playerMatch.WalkDistance, &playerMatch.TimeSurvived, &playerMatch.WinPlace)
		playerMatches = append(playerMatches, playerMatch)
		checkErr(err)

	}

	defer rows.Close()
	return playerMatches
}

func (m MySQL) CreatePlayerMatch(playerMatch PlayerMatch) PlayerMatch {
	playerMatch.ID = xid.New().String()
	stmt, err := m.db.Prepare("insert into PlayerMatch(id, matchID, playerID, DBNOs, assists, damageDealt, headshotKills, longestKill, kills, revives, rideDistance, swimDistance, walkDistance, timeSurvived, winPlace) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	checkErr(err)

	_, err = stmt.Exec(playerMatch.ID, playerMatch.MatchID, playerMatch.PlayerID, playerMatch.DBNOs, playerMatch.Assists, playerMatch.DamageDealt, playerMatch.HeadshotKills, playerMatch.LongestKill, playerMatch.Kills, playerMatch.Revives, playerMatch.RideDistance, playerMatch.SwimDistance, playerMatch.WalkDistance, playerMatch.TimeSurvived, playerMatch.WinPlace)
	checkErr(err)

	return playerMatch
}

func (m MySQL) GetPlayerMatch(playerMatch PlayerMatch) PlayerMatch {
	rows, err := m.db.Query("select * from PlayerMatch where playerID = ? and matchID = ?;", playerMatch.PlayerID, playerMatch.MatchID)

	checkErr(err)

	var p PlayerMatch
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.MatchID, &p.PlayerID, &p.DBNOs, &p.Assists, &p.DamageDealt, &p.HeadshotKills, &p.LongestKill, &p.Kills, &p.Revives, &p.RideDistance, &p.SwimDistance, &p.WalkDistance, &p.TimeSurvived, &p.WinPlace)
		checkErr(err)

	}

	defer rows.Close()
	return p
}

func (m MySQL) UpdatePlayerMatch(playerMatch PlayerMatch) PlayerMatch {
	stmt, err := m.db.Prepare("update PlayerMatch set matchID=?, playerID=?, DBNOs=?, assists=?, damageDealt=?, headshotKills=?, longestKill=?, kills=?, revives=?, rideDistance=?, swimDistance=?, walkDistance=?, timeSurvived=?, winPlace=? where matchID=? and playerID=?")
	checkErr(err)

	_, err = stmt.Exec(playerMatch.MatchID, playerMatch.PlayerID, playerMatch.DBNOs, playerMatch.Assists, playerMatch.DamageDealt, playerMatch.HeadshotKills, playerMatch.LongestKill, playerMatch.Kills, playerMatch.Revives, playerMatch.RideDistance, playerMatch.SwimDistance, playerMatch.WalkDistance, playerMatch.TimeSurvived, playerMatch.WinPlace, playerMatch.MatchID, playerMatch.PlayerID)
	checkErr(err)

	return playerMatch
}

func (m MySQL) DeletePlayerMatch(playerMatch PlayerMatch) PlayerMatch {
	stmt, err := m.db.Prepare("delete from PlayerMatch where playerID=? and matchID=?;")
	checkErr(err)

	_, err = stmt.Exec(playerMatch.PlayerID, playerMatch.MatchID)
	checkErr(err)

	return playerMatch
}

func (m MySQL) GetCupMatches(cupMatch CupMatch) []CupMatch {
	var (
		cupMatches []CupMatch
	)

	rows, err := m.db.Query("select * from CupMatch where cupID = ?;", cupMatch.CupID)

	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&cupMatch.ID, &cupMatch.CupID, &cupMatch.MatchID)
		cupMatches = append(cupMatches, cupMatch)
		checkErr(err)

	}

	defer rows.Close()
	return cupMatches
}

func (m MySQL) CreateCupMatch(cupMatch CupMatch) CupMatch {
	cupMatch.ID = xid.New().String()
	stmt, err := m.db.Prepare("insert into CupMatch(id, matchID, cupID) VALUES (?, ?, ?)")
	checkErr(err)

	_, err = stmt.Exec(cupMatch.ID, cupMatch.MatchID, cupMatch.CupID)
	checkErr(err)

	return cupMatch
}

func (m MySQL) GetCupMatch(cupMatch CupMatch) CupMatch {
	rows, err := m.db.Query("select * from CupMatch where cupID = ? and matchID = ?;", cupMatch.CupID, cupMatch.MatchID)

	checkErr(err)

	var c CupMatch
	for rows.Next() {
		err = rows.Scan(&c.ID, &c.CupID, &c.MatchID)
		checkErr(err)

	}

	defer rows.Close()
	return c
}

func (m MySQL) UpdateCupMatch(cupMatch CupMatch) CupMatch {
	stmt, err := m.db.Prepare("update CupMatch set matchID=?, cupID=? where id=?")
	checkErr(err)

	_, err = stmt.Exec(cupMatch.MatchID, cupMatch.CupID, cupMatch.ID)
	checkErr(err)

	return cupMatch
}

func (m MySQL) DeleteCupMatch(cupMatch CupMatch) CupMatch {
	stmt, err := m.db.Prepare("delete from CupMatch where matchID=?, cupID=?")
	checkErr(err)

	_, err = stmt.Exec(cupMatch.MatchID, cupMatch.CupID)
	checkErr(err)

	return cupMatch
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
