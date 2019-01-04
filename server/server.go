package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bavaz1/go-pubg/database"
	"github.com/go-chi/chi"
	"github.com/gorilla/mux"
)

type Server struct {
	client  *http.Client
	address string
	storage database.Storage
}

func (s *Server) Listen() {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/players", func(r chi.Router) {
			r.Get("/", s.getPlayers)
			r.Post("/", s.createPlayer)
			r.Get("/{id}", s.getPlayer)
			r.Put("/{id}", s.putPlayer)
			r.Delete("/{id}", s.deletePlayer)
			r.Route("/{playerId}/matches", func(r chi.Router) {
				r.Get("/", s.getPlayerMatches)
				r.Post("/", s.createPlayerMatch)
				r.Get("/{id}", s.getPlayerMatch)
				r.Put("/{id}", s.putPlayerMatch)
				r.Delete("/{id}", s.deletePlayerMatch)
			})
		})
		r.Route("/match", func(r chi.Router) {
			r.Get("/", s.getMatches)
			r.Post("/", s.createMatch)
			r.Get("/{id}", s.getMatch)
			r.Put("/{id}", s.putMatch)
			r.Delete("/{id}", s.deleteMatch)
		})
		r.Route("/cups", func(r chi.Router) {
			r.Get("/", s.getCups)
			r.Post("/", s.createCup)
			r.Get("/{id}", s.getCup)
			r.Put("/{id}", s.putCup)
			r.Delete("/{id}", s.deleteCup)
			r.Route("/{cupId}/matches", func(r chi.Router) {
				r.Get("/", s.getCupMatches)
				r.Post("/", s.createCupMatch)
				r.Get("/{id}", s.getCupMatch)
				r.Put("/{id}", s.putCupMatch)
				r.Delete("/{id}", s.deleteCupMatch)
			})
		})
	})

	http.ListenAndServe(s.address, r)
}

func (s *Server) getPlayers(w http.ResponseWriter, r *http.Request) {
	p := s.storage.GetPlayers()

	response, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) createPlayer(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var player database.Player
	err = json.Unmarshal(b, &player)
	if err != nil {
		panic(err)
	}

	p := s.storage.CreatePlayer(player)

	response, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) getPlayer(w http.ResponseWriter, r *http.Request) {
	var player database.Player

	params := mux.Vars(r)
	player.ID = params["id"]

	p := s.storage.GetPlayer(player)

	response, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) putPlayer(w http.ResponseWriter, r *http.Request) {
	var player database.Player

	params := mux.Vars(r)
	player.ID = params["id"]

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	err = json.Unmarshal(b, &player)
	if err != nil {
		panic(err)
	}

	p := s.storage.UpdatePlayer(player)

	response, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) deletePlayer(w http.ResponseWriter, r *http.Request) {
	var player database.Player

	params := mux.Vars(r)
	player.ID = params["id"]

	p := s.storage.DeletePlayer(player)

	response, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) getMatches(w http.ResponseWriter, r *http.Request) {
	m := s.storage.GetMatches()

	response, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) createMatch(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var match database.Match
	err = json.Unmarshal(b, &match)
	if err != nil {
		panic(err)
	}

	m := s.storage.CreateMatch(match)

	response, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) getMatch(w http.ResponseWriter, r *http.Request) {
	var match database.Match

	params := mux.Vars(r)
	match.ID = params["id"]

	m := s.storage.GetMatch(match)

	response, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) putMatch(w http.ResponseWriter, r *http.Request) {
	var match database.Match

	params := mux.Vars(r)
	match.ID = params["id"]

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	err = json.Unmarshal(b, &match)
	if err != nil {
		panic(err)
	}

	m := s.storage.UpdateMatch(match)

	response, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) deleteMatch(w http.ResponseWriter, r *http.Request) {
	var match database.Match

	params := mux.Vars(r)
	match.ID = params["id"]

	m := s.storage.DeleteMatch(match)

	response, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) getCups(w http.ResponseWriter, r *http.Request) {
	c := s.storage.GetCups()

	response, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) createCup(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var cup database.Cup
	err = json.Unmarshal(b, &cup)
	if err != nil {
		panic(err)
	}

	c := s.storage.CreateCup(cup)

	response, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) getCup(w http.ResponseWriter, r *http.Request) {
	var cup database.Cup

	params := mux.Vars(r)
	cup.ID = params["id"]

	c := s.storage.GetCup(cup)

	response, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) putCup(w http.ResponseWriter, r *http.Request) {
	var cup database.Cup

	params := mux.Vars(r)
	cup.ID = params["id"]

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	err = json.Unmarshal(b, &cup)
	if err != nil {
		panic(err)
	}

	c := s.storage.UpdateCup(cup)

	response, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) deleteCup(w http.ResponseWriter, r *http.Request) {
	var cup database.Cup

	params := mux.Vars(r)
	cup.ID = params["id"]

	c := s.storage.DeleteCup(cup)

	response, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) getPlayerMatches(w http.ResponseWriter, r *http.Request) {
	var playerMatch database.PlayerMatch

	params := mux.Vars(r)
	playerMatch.PlayerID = params["playerId"]

	p := s.storage.GetPlayerMatches(playerMatch)

	response, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) createPlayerMatch(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var playerMatch database.PlayerMatch
	err = json.Unmarshal(b, &playerMatch)
	if err != nil {
		panic(err)
	}

	p := s.storage.CreatePlayerMatch(playerMatch)

	response, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) getPlayerMatch(w http.ResponseWriter, r *http.Request) {
	var playerMatch database.PlayerMatch

	params := mux.Vars(r)
	playerMatch.MatchID = params["id"]
	playerMatch.PlayerID = params["playerId"]

	p := s.storage.GetPlayerMatch(playerMatch)

	response, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) putPlayerMatch(w http.ResponseWriter, r *http.Request) {
	var playerMatch database.PlayerMatch

	params := mux.Vars(r)
	playerMatch.MatchID = params["id"]
	playerMatch.PlayerID = params["playerId"]

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	err = json.Unmarshal(b, &playerMatch)
	if err != nil {
		panic(err)
	}

	p := s.storage.UpdatePlayerMatch(playerMatch)

	response, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) deletePlayerMatch(w http.ResponseWriter, r *http.Request) {
	var playerMatch database.PlayerMatch

	params := mux.Vars(r)
	playerMatch.MatchID = params["id"]
	playerMatch.PlayerID = params["playerId"]

	p := s.storage.DeletePlayerMatch(playerMatch)

	response, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) getCupMatches(w http.ResponseWriter, r *http.Request) {
	var cupMatch database.CupMatch

	params := mux.Vars(r)
	cupMatch.CupID = params["cupId"]

	c := s.storage.GetCupMatch(cupMatch)

	response, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) createCupMatch(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	var cupMatch database.CupMatch
	err = json.Unmarshal(b, &cupMatch)
	if err != nil {
		panic(err)
	}

	p := s.storage.CreateCupMatch(cupMatch)

	response, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) getCupMatch(w http.ResponseWriter, r *http.Request) {
	var cupMatch database.CupMatch

	params := mux.Vars(r)
	cupMatch.CupID = params["cupId"]
	cupMatch.MatchID = params["id"]

	c := s.storage.GetCupMatch(cupMatch)

	response, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) putCupMatch(w http.ResponseWriter, r *http.Request) {
	var cupMatch database.CupMatch

	params := mux.Vars(r)
	cupMatch.CupID = params["cupId"]
	cupMatch.MatchID = params["id"]

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	err = json.Unmarshal(b, &cupMatch)
	if err != nil {
		panic(err)
	}

	c := s.storage.UpdateCupMatch(cupMatch)

	response, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (s *Server) deleteCupMatch(w http.ResponseWriter, r *http.Request) {
	var cupMatch database.CupMatch

	params := mux.Vars(r)
	cupMatch.CupID = params["cupId"]
	cupMatch.MatchID = params["id"]

	c := s.storage.DeletePlayerMatch(cupMatch)

	response, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
