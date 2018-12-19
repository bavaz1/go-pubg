package database

import (
	"testing"
	"time"
)

func TestPlayers(t *testing.T) {
	m := New()
	p := m.CreatePlayer(Player{
		Name:       "Gyozo",
		Email:      "test2@gmail.com",
		InGameName: "bGYOZTES",
	})
	if p.Name != "Gyozo" {
		t.Errorf("Name should be 'Gyozo', instead of %s", p.Name)
	}
	if p.Email != "test2@gmail.com" {
		t.Errorf("Email should be 'test2@gmail.com', instead of %s", p.Email)
	}
	if p.InGameName != "bGYOZTES" {
		t.Errorf("InGameName should be 'bGYOZTES', instead of %s", p.InGameName)
	}
	// t.Log(p)

	p2 := m.GetPlayer(p)
	if p2.Name != "Gyozo" {
		t.Errorf("Name should be 'Gyozo', instead of %s", p2.Name)
	}
	if p2.Email != "test2@gmail.com" {
		t.Errorf("Email should be 'test2@gmail.com', instead of %s", p2.Email)
	}
	if p2.InGameName != "bGYOZTES" {
		t.Errorf("InGameName should be 'bGYOZTES', instead of %s", p2.InGameName)
	}
	// t.Log(p2)

	p3 := m.UpdatePlayer(Player{
		ID:         p2.ID,
		Name:       "ChangedName",
		Email:      p2.Email,
		InGameName: "bro",
	})
	if p3.Name != "ChangedName" {
		t.Errorf("Name should be 'ChangedName', instead of %s", p3.Name)
	}
	if p3.Email != "test2@gmail.com" {
		t.Errorf("Email should be 'test2@gmail.com', instead of %s", p3.Email)
	}
	if p3.InGameName != "bro" {
		t.Errorf("InGameName should be 'bro', instead of %s", p3.InGameName)
	}
	// t.Log(p3)

	p4 := m.GetPlayer(Player{
		ID: p2.ID,
	})
	if p4.Name != "ChangedName" {
		t.Errorf("Name should be 'ChangedName', instead of %s", p4.Name)
	}
	if p4.Email != "test2@gmail.com" {
		t.Errorf("Email should be 'test2@gmail.com', instead of %s", p4.Email)
	}
	if p4.InGameName != "bro" {
		t.Errorf("InGameName should be 'bro', instead of %s", p4.InGameName)
	}
	// t.Log(p4)

	p5 := m.DeletePlayer(Player{
		ID: p2.ID,
	})
	if p5.Name == "ChangedName" {
		t.Errorf("Name should be %s, instead of 'ChangedName'", p4.Name)
	}
	if p5.Email == "test2@gmail.com" {
		t.Errorf("Email should be %s, instead of 'test2@gmail.com'", p4.Email)
	}
	if p5.InGameName == "bro" {
		t.Errorf("InGameName should be %s, instead of 'bro'", p4.InGameName)
	}
	// t.Log(p5)

	p6 := m.GetPlayers()
	if len(p6) == 0 {
		t.Errorf("Players array length should be not 0")
	}
	// t.Log(p6)
}

func TestCups(t *testing.T) {
	m := New()
	c := m.CreateCup(Cup{
		BeginAt:  time.Date(2019, time.Month(1), 15, 10, 30, 0, 0, time.UTC),
		GameMode: "solo",
	})
	if c.BeginAt != time.Date(2019, time.Month(1), 15, 10, 30, 0, 0, time.UTC) {
		t.Errorf("BeginAt should be "+time.Date(2019, time.Month(1), 15, 10, 30, 0, 0, time.UTC).String()+", instead of %s", c.BeginAt)
	}
	if c.GameMode != "solo" {
		t.Errorf("GameMode should be 'solo', instead of %s", c.GameMode)
	}
	// t.Log(c)

	c2 := m.GetCup(Cup{
		ID: c.ID,
	})
	if c2.BeginAt != time.Date(2019, time.Month(1), 15, 10, 30, 0, 0, time.UTC) {
		t.Errorf("BeginAt should be "+time.Date(2019, time.Month(1), 15, 10, 30, 0, 0, time.UTC).String()+", instead of %s", c2.BeginAt)
	}
	if c2.GameMode != "solo" {
		t.Errorf("GameMode should be 'solo', instead of %s", c2.GameMode)
	}
	// t.Log(c2)

	p := m.CreatePlayer(Player{
		Name:       "Gyozo",
		Email:      "test2@gmail.com",
		InGameName: "bGYOZTES",
	})
	if p.Name != "Gyozo" {
		t.Errorf("Name should be 'Gyozo', instead of %s", p.Name)
	}
	if p.Email != "test2@gmail.com" {
		t.Errorf("Email should be 'test2@gmail.com', instead of %s", p.Email)
	}
	if p.InGameName != "bGYOZTES" {
		t.Errorf("InGameName should be 'bGYOZTES', instead of %s", p.InGameName)
	}
	// t.Log(p)
	c3 := m.UpdateCup(Cup{
		ID:       c.ID,
		BeginAt:  c.BeginAt,
		Winner:   p.ID,
		GameMode: c.GameMode,
	})
	if c3.BeginAt != time.Date(2019, time.Month(1), 15, 10, 30, 0, 0, time.UTC) {
		t.Errorf("BeginAt should be %s, instead of '%s'", time.Date(2019, time.Month(1), 15, 10, 30, 0, 0, time.UTC).String(), c3.BeginAt)
	}
	if c3.Winner != p.ID {
		t.Errorf("Winner should be '%s', instead of %s", p.ID, c3.Winner)
	}
	if c2.GameMode != "solo" {
		t.Errorf("GameMode should be 'solo', instead of %s", c2.GameMode)
	}
	// t.Log(c3)

	c4 := m.DeleteCup(Cup{
		ID: c3.ID,
	})
	if c4.BeginAt == time.Date(2019, time.Month(1), 15, 10, 30, 0, 0, time.UTC) {
		t.Errorf("BeginAt should be 'NULL', instead of '%s'", c3.BeginAt)
	}
	if c4.Winner == p.ID {
		t.Errorf("Winner should be 'NULL', instead of '%s'", c3.ID)
	}
	if c4.GameMode == "solo" {
		t.Errorf("GameMode should be 'NULL', instead of %s", c3.GameMode)
	}
	// t.Log(c4)

	c5 := m.GetCups()
	if len(c5) == 0 {
		t.Errorf("Cups array length should be not 0")
	}
	// t.Log(c5)
}

func TestMatch(t *testing.T) {
	m := New()
	begin := time.Now()
	match := m.CreateMatch(Match{
		Duration: 1890,
		MapName:  "erangel",
		IsCustom: true,
		BeginAt:  begin,
	})
	if match.Duration != 1890 {
		t.Errorf("Duration should be 1890, instead of %d", match.Duration)
	}
	if match.MapName != "erangel" {
		t.Errorf("MapName should be 'erangel', instead of %s", match.MapName)
	}
	if match.IsCustom != true {
		t.Errorf("IsCustom should be true, instead of %t", match.IsCustom)
	}
	if match.BeginAt != begin {
		t.Errorf("BeginAt should be %s, instead of %s", begin.String(), match.BeginAt.String())
	}
	// t.Log(match)

	match2 := m.GetMatch(Match{
		ID: match.ID,
	})
	if match2.Duration != 1890 {
		t.Errorf("Duration should be 1890, instead of %d", match2.Duration)
	}
	if match2.MapName != "erangel" {
		t.Errorf("MapName should be 'erangel', instead of %s", match2.MapName)
	}
	if match2.IsCustom != true {
		t.Errorf("IsCustom should be true, instead of %t", match2.IsCustom)
	}
	/* if match2.BeginAt != match.BeginAt {
		t.Errorf("BeginAt should be %s, instead of %s", match.BeginAt.String(), match2.BeginAt.String())
	} */
	// t.Log(match2)

	match3 := m.UpdateMatch(Match{
		ID:       match2.ID,
		Duration: match2.Duration,
		MapName:  "sanhok",
		IsCustom: match2.IsCustom,
		BeginAt:  match2.BeginAt,
	})
	if match3.Duration != match2.Duration {
		t.Errorf("Duration should be %d, instead of %d", match2.Duration, match3.Duration)
	}
	if match3.MapName != "sanhok" {
		t.Errorf("MapName should be 'sanhok', instead of %s", match3.MapName)
	}
	if match3.IsCustom != match.IsCustom {
		t.Errorf("IsCustom should be %t, instead of %t", match2.IsCustom, match3.IsCustom)
	}
	// t.Log(match3)

	match4 := m.DeleteMatch(Match{
		ID: match3.ID,
	})
	if match4.Duration == 1890 {
		t.Errorf("Duration should be 'NULL', instead of '%d'", match4.Duration)
	}
	if match4.MapName == "sanhok" {
		t.Errorf("MapName should be 'NULL', instead of '%s'", match4.MapName)
	}
	if match4.IsCustom == true {
		t.Errorf("IsCustom should be 'NULL', instead of %t", match4.IsCustom)
	}
	// t.Log(match4)

	match5 := m.GetMatches()
	if len(match5) == 0 {
		t.Errorf("Matches array length should be not 0")
	}
	// t.Log(match5)
}

func TestPlayerMatch(t *testing.T) {
	m := New()

	player := m.CreatePlayer(Player{
		Name:       "Gyozo",
		Email:      "test2@gmail.com",
		InGameName: "bGYOZTES",
	})
	if player.Name != "Gyozo" {
		t.Errorf("Name should be 'Gyozo', instead of %s", player.Name)
	}
	if player.Email != "test2@gmail.com" {
		t.Errorf("Email should be 'test2@gmail.com', instead of %s", player.Email)
	}
	if player.InGameName != "bGYOZTES" {
		t.Errorf("InGameName should be 'bGYOZTES', instead of %s", player.InGameName)
	}

	begin := time.Now()
	match := m.CreateMatch(Match{
		Duration: 1890,
		MapName:  "erangel",
		IsCustom: true,
		BeginAt:  begin,
	})
	if match.Duration != 1890 {
		t.Errorf("Duration should be 1890, instead of %d", match.Duration)
	}
	if match.MapName != "erangel" {
		t.Errorf("MapName should be 'erangel', instead of %s", match.MapName)
	}
	if match.IsCustom != true {
		t.Errorf("IsCustom should be true, instead of %t", match.IsCustom)
	}
	if match.BeginAt != begin {
		t.Errorf("BeginAt should be %s, instead of %s", begin.String(), match.BeginAt.String())
	}

	p := m.CreatePlayerMatch(PlayerMatch{
		MatchID:       match.ID,
		PlayerID:      player.ID,
		DBNOs:         3,
		Assists:       3,
		DamageDealt:   452.81,
		HeadshotKills: 2,
		LongestKill:   105.14,
		Kills:         3,
		Revives:       3,
		RideDistance:  321,
		SwimDistance:  321.21,
		WalkDistance:  321.21,
		TimeSurvived:  321.21,
		WinPlace:      21,
	})
	if p.DBNOs != 3 {
		t.Errorf("DBNOs should be 3, instead of %d", p.DBNOs)
	}
	if p.Assists != 3 {
		t.Errorf("Assists should be 3, instead of %d", p.Assists)
	}
	if p.DamageDealt != 452.81 {
		t.Errorf("DamageDealt should be 452.81, instead of %f", p.DamageDealt)
	}
	if p.HeadshotKills != 2 {
		t.Errorf("HeadshotKills should be 2, instead of %d", p.HeadshotKills)
	}
	if p.LongestKill != 105.14 {
		t.Errorf("LongestKill should be 105.14, instead of %f", p.LongestKill)
	}
	if p.Kills != 3 {
		t.Errorf("Kills should be 3, instead of %d", p.Kills)
	}
	if p.Revives != 3 {
		t.Errorf("Revives should be 3, instead of %d", p.Revives)
	}
	if p.RideDistance != 321 {
		t.Errorf("RideDistance should be 321, instead of %f", p.RideDistance)
	}
	if p.SwimDistance != 321.21 {
		t.Errorf("SwimDistance should be 321.21, instead of %f", p.SwimDistance)
	}
	if p.WalkDistance != 321.21 {
		t.Errorf("WalkDistance should be 321.21, instead of %f", p.WalkDistance)
	}
	if p.TimeSurvived != 321.21 {
		t.Errorf("TimeSurvived should be 321.21, instead of %f", p.TimeSurvived)
	}
	if p.WinPlace != 21 {
		t.Errorf("WinPlace should be 21, instead of %d", p.WinPlace)
	}
	// t.Log(p)

	p2 := m.GetPlayerMatch(PlayerMatch{
		ID: p.ID,
	})
	if p2.DBNOs != 3 {
		t.Errorf("DBNOs should be 3, instead of %d", p2.DBNOs)
	}
	if p2.Assists != 3 {
		t.Errorf("Assists should be 3, instead of %d", p2.Assists)
	}
	if p2.DamageDealt != 452.81 {
		t.Errorf("DamageDealt should be 452.81, instead of %f", p2.DamageDealt)
	}
	if p2.HeadshotKills != 2 {
		t.Errorf("HeadshotKills should be 2, instead of %d", p2.HeadshotKills)
	}
	if p2.LongestKill != 105.14 {
		t.Errorf("LongestKill should be 105.14, instead of %f", p2.LongestKill)
	}
	if p2.Kills != 3 {
		t.Errorf("Kills should be 3, instead of %d", p2.Kills)
	}
	if p2.Revives != 3 {
		t.Errorf("Revives should be 3, instead of %d", p2.Revives)
	}
	if p2.RideDistance != 321 {
		t.Errorf("RideDistance should be 321, instead of %f", p2.RideDistance)
	}
	if p2.SwimDistance != 321.21 {
		t.Errorf("SwimDistance should be 321.21, instead of %f", p2.SwimDistance)
	}
	if p2.WalkDistance != 321.21 {
		t.Errorf("WalkDistance should be 321.21, instead of %f", p2.WalkDistance)
	}
	if p2.TimeSurvived != 321.21 {
		t.Errorf("TimeSurvived should be 321.21, instead of %f", p2.TimeSurvived)
	}
	if p2.WinPlace != 21 {
		t.Errorf("WinPlace should be 21, instead of %d", p2.WinPlace)
	}
	// t.Log(p2)

	p3 := m.UpdatePlayerMatch(PlayerMatch{
		ID:            p.ID,
		MatchID:       match.ID,
		PlayerID:      player.ID,
		DBNOs:         4,
		Assists:       3,
		DamageDealt:   300.81,
		HeadshotKills: 2,
		LongestKill:   81.14,
		Kills:         2,
		Revives:       3,
		RideDistance:  103,
		SwimDistance:  324.21,
		WalkDistance:  321.21,
		TimeSurvived:  321.21,
		WinPlace:      38,
	})
	if p3.DBNOs != 4 {
		t.Errorf("DBNOs should be 4, instead of %d", p3.DBNOs)
	}
	if p3.Assists != 3 {
		t.Errorf("Assists should be 3, instead of %d", p3.Assists)
	}
	if p3.DamageDealt != 300.81 {
		t.Errorf("DamageDealt should be 300.81, instead of %f", p3.DamageDealt)
	}
	if p3.HeadshotKills != 2 {
		t.Errorf("HeadshotKills should be 2, instead of %d", p3.HeadshotKills)
	}
	if p3.LongestKill != 81.14 {
		t.Errorf("LongestKill should be 81.14, instead of %f", p3.LongestKill)
	}
	if p3.Kills != 2 {
		t.Errorf("Kills should be 2, instead of %d", p3.Kills)
	}
	if p3.Revives != 3 {
		t.Errorf("Revives should be 3, instead of %d", p3.Revives)
	}
	if p3.RideDistance != 103 {
		t.Errorf("RideDistance should be 103, instead of %f", p3.RideDistance)
	}
	if p3.SwimDistance != 324.21 {
		t.Errorf("SwimDistance should be 324.21, instead of %f", p3.SwimDistance)
	}
	if p3.WalkDistance != 321.21 {
		t.Errorf("WalkDistance should be 321.21, instead of %f", p3.WalkDistance)
	}
	if p3.TimeSurvived != 321.21 {
		t.Errorf("TimeSurvived should be 321.21, instead of %f", p3.TimeSurvived)
	}
	if p3.WinPlace != 38 {
		t.Errorf("WinPlace should be 38, instead of %d", p3.WinPlace)
	}
	// t.Log(p3)

	p4 := m.DeletePlayerMatch(PlayerMatch{
		ID: p.ID,
	})
	if p4.DBNOs == 4 {
		t.Errorf("DBNOs should be NULL, instead of %d", p4.DBNOs)
	}
	if p4.Assists == 3 {
		t.Errorf("Assists should be NULL, instead of %d", p4.Assists)
	}
	if p4.DamageDealt == 300.81 {
		t.Errorf("DamageDealt should be NULL, instead of %f", p4.DamageDealt)
	}
	if p4.HeadshotKills == 2 {
		t.Errorf("HeadshotKills should be NULL, instead of %d", p4.HeadshotKills)
	}
	if p4.LongestKill == 81.14 {
		t.Errorf("LongestKill should be NULL, instead of %f", p4.LongestKill)
	}
	if p4.Kills == 2 {
		t.Errorf("Kills should be NULL, instead of %d", p4.Kills)
	}
	if p4.Revives == 3 {
		t.Errorf("Revives should be NULL, instead of %d", p4.Revives)
	}
	if p4.RideDistance == 103 {
		t.Errorf("RideDistance should be NULL, instead of %f", p4.RideDistance)
	}
	if p4.SwimDistance == 324.21 {
		t.Errorf("SwimDistance should be NULL, instead of %f", p4.SwimDistance)
	}
	if p4.WalkDistance == 321.21 {
		t.Errorf("WalkDistance should be NULL, instead of %f", p4.WalkDistance)
	}
	if p4.TimeSurvived == 321.21 {
		t.Errorf("TimeSurvived should be NULL, instead of %f", p4.TimeSurvived)
	}
	if p4.WinPlace == 38 {
		t.Errorf("WinPlace should be NULL, instead of %d", p4.WinPlace)
	}
	// t.Log(p4)

	p5 := m.GetPlayerMatches()
	if len(p5) == 0 {
		t.Errorf("PlayerMatches array length should be not 0")
	}
	// t.Log(p5)
}

func TestCupMatch(t *testing.T) {
	m := New()
	cup := m.CreateCup(Cup{
		BeginAt:  time.Date(2019, time.Month(1), 15, 10, 30, 0, 0, time.UTC),
		GameMode: "solo",
	})
	if cup.BeginAt != time.Date(2019, time.Month(1), 15, 10, 30, 0, 0, time.UTC) {
		t.Errorf("BeginAt should be "+time.Date(2019, time.Month(1), 15, 10, 30, 0, 0, time.UTC).String()+", instead of %s", cup.BeginAt)
	}
	if cup.GameMode != "solo" {
		t.Errorf("GameMode should be 'solo', instead of %s", cup.GameMode)
	}
	// t.Log(cup)

	cup2 := m.CreateCup(Cup{
		BeginAt:  time.Date(2019, time.Month(1), 15, 10, 30, 0, 0, time.UTC),
		GameMode: "duo",
	})
	if cup2.BeginAt != time.Date(2019, time.Month(1), 15, 10, 30, 0, 0, time.UTC) {
		t.Errorf("BeginAt should be "+time.Date(2019, time.Month(1), 15, 10, 30, 0, 0, time.UTC).String()+", instead of %s", cup2.BeginAt)
	}
	if cup2.GameMode != "duo" {
		t.Errorf("GameMode should be 'duo', instead of %s", cup2.GameMode)
	}
	// t.Log(cup2)

	begin := time.Now()
	match := m.CreateMatch(Match{
		Duration: 1890,
		MapName:  "erangel",
		IsCustom: true,
		BeginAt:  begin,
	})
	if match.Duration != 1890 {
		t.Errorf("Duration should be 1890, instead of %d", match.Duration)
	}
	if match.MapName != "erangel" {
		t.Errorf("MapName should be 'erangel', instead of %s", match.MapName)
	}
	if match.IsCustom != true {
		t.Errorf("IsCustom should be true, instead of %t", match.IsCustom)
	}
	if match.BeginAt != begin {
		t.Errorf("BeginAt should be %s, instead of %s", begin.String(), match.BeginAt.String())
	}
	// t.Log(match)

	c := m.CreateCupMatch(CupMatch{
		CupID:   cup.ID,
		MatchID: match.ID,
	})
	if c.CupID != cup.ID {
		t.Errorf("Cup ID should be %s, instead of %s", cup.ID, c.CupID)
	}
	if c.MatchID != match.ID {
		t.Errorf("Email should be %s, instead of %s", match.ID, c.MatchID)
	}
	// t.Log(c)

	c2 := m.GetCupMatch(CupMatch{
		ID: c.ID,
	})
	if c2.CupID != c.CupID {
		t.Errorf("CupID should be %s, instead of %s", c.CupID, c2.CupID)
	}
	if c2.MatchID != c.MatchID {
		t.Errorf("MatchID should be %s, instead of %s", c.MatchID, c2.MatchID)
	}
	// t.Log(c2)

	c3 := m.UpdateCupMatch(CupMatch{
		ID:      c.ID,
		CupID:   cup2.ID,
		MatchID: c.MatchID,
	})
	if c3.CupID != cup2.ID {
		t.Errorf("CupID should be %s, instead of %s", cup2.ID, c3.CupID)
	}
	if c3.MatchID != c.MatchID {
		t.Errorf("MatchID should be %s, instead of %s", c.MatchID, c3.MatchID)
	}
	// t.Log(c3)

	c4 := m.DeleteCupMatch(CupMatch{
		ID: c.ID,
	})
	if c4.CupID == c.CupID {
		t.Errorf("CupID should be NULL, instead of %s", c4.CupID)
	}
	if c4.MatchID == c.MatchID {
		t.Errorf("MatchID should be NULL, instead of %s", c4.MatchID)
	}
	// t.Log(c4)

	c5 := m.GetCupMatches()
	if len(c5) == 0 {
		t.Errorf("CupMatches array length should be not 0")
	}
	// t.Log(c5)
}
