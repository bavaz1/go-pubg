package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Base struct {
	Data     []Data      `json:"data"`
	Links    Links       `json:"links"`
	Meta     interface{} `json:"meta"`
	Included []Data      `json:"included"`
}

type Data struct {
	Type          string        `json:"type"`
	ID            string        `json:"id"`
	Attributes    Attributes    `json:"attributes"`
	Relationships Relationships `json:"relationships"`
	Links         Links         `json:"links"`
}

type Attributes struct {
	CreatedAt     string      `json:"createdAt"`
	UpdatedAt     string      `json:"updatedAt"`
	PatchVersion  string      `json:"patchVersion"`
	Name          string      `json:"name"`
	Stats         Stats       `json:"stats"`
	TitleID       string      `json:"titleId"`
	ShardID       string      `json:"shardId"`
	IsCustomMatch bool        `json:"isCustomMatch"`
	SeasonState   string      `json:"seasonState"`
	Tags          interface{} `json:"tags"`
	MapName       string      `json:"mapName"`
	GameMode      string      `json:"gameMode"`
	Duration      int         `json:"duration"`
	Actor         string      `json:"actor"`
}

type Relationships struct {
	Assets  Assets  `json:"assets"`
	Matches Matches `json:"matches"`
	Rosters Rosters `json:"rosters"`
}

type Matches struct {
	Data []SubData `json:"data"`
}

type SubData struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type Links struct {
	Self   string `json:"self"`
	Schema string `json:"schema"`
}

type Rosters struct {
	Data []SubData `json:"data"`
}

type Assets struct {
	Data []SubData `json:"data"`
}

type Stats struct {
	DBNOs           int     `json:"DBNOs"`
	Assists         int     `json:"assists"`
	Boosts          int     `json:"boosts"`
	DamageDealt     float32 `json:"damageDealt"`
	DeathType       string  `json:"deathType"`
	HeadshotKills   int     `json:"headshotKills"`
	Heals           int     `json:"heals"`
	KillPlace       int     `json:"killPlace"`
	KillPoints      int     `json:"killPoints"`
	KillPointsDelta float32 `json:"killPointsDelta"`
	KillStreaks     int     `json:"killStreaks"`
	Kills           int     `json:"kills"`
	LastKillPoints  int     `json:"lastKillPoints"`
	LastWinPoints   int     `json:"lastWinPoints"`
	LongestKill     int     `json:"longestKill"`
	MostDamage      int     `json:"mostDamage"`
	Name            string  `json:"name"`
	PlayerID        string  `json:"playerId"`
	RankPoints      int     `json:"rankPoints"`
	Revives         int     `json:"revives"`
	RideDistance    float32 `json:"rideDistance"`
	SwimDistance    int     `json:"swimDistance"`
	TeamKills       int     `json:"teamKills"`
	TimeSurvived    float64 `json:"timeSurvived"`
	VehicleDestroys int     `json:"vehicleDestroys"`
	WalkDistance    float64 `json:"walkDistance"`
	WeaponsAcquired int     `json:"weaponsAcquired"`
	WinPlace        int     `json:"winPlace"`
	WinPoints       int     `json:"winPoints"`
	WinPointsDelta  float32 `json:"winPointsDelta"`
}

const apiHost = "https://api.pubg.com"
const apiEndpoint = "shards"
const apiShard = "steam"
const apiFilter = "players?filter[playerNames]="
const apiMatches = "matches"

func GetPlayer(ctx context.Context, playerName string, client *http.Client) (Base, error) {
	s := []string{apiHost, apiEndpoint, apiShard, apiFilter}
	url := strings.Join(s, "/") + playerName

	jsonFile, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	apiKey := result["pubgApiKey"]

	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Base{}, err
	}
	r.Header.Set("Accept", "application/vnd.api+json")
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	r = r.WithContext(ctx)

	resp, err := client.Do(r)
	if err != nil {
		return Base{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Base{}, err
	}

	if resp.StatusCode == 200 {
		var result Base
		err = json.Unmarshal(body, &result)
		return result, err
	}

	return Base{}, fmt.Errorf("Undefined status code: %s", string(body))
}

func GetMatch(ctx context.Context, matchID string, client *http.Client) (Base, error) {
	s := []string{apiHost, apiEndpoint, apiShard, apiMatches, matchID}
	url := strings.Join(s, "/")

	jsonFile, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	apiKey := result["pubgApiKey"]

	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Base{}, err
	}
	r.Header.Set("Accept", "application/vnd.api+json")
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	r = r.WithContext(ctx)

	resp, err := client.Do(r)
	if err != nil {
		return Base{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Base{}, err
	}

	if resp.StatusCode == 200 {
		var result Base
		err = json.Unmarshal(body, &result)
		fmt.Println(result)
		return result, err
	}

	return Base{}, fmt.Errorf("Undefined status code: %s", string(body))
}
