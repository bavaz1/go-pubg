package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bavaz1/go-pubg/database"
	"github.com/go-chi/chi"
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
			r.Post("/", s.createPlayer)
		})
		r.Route("/match", func(r chi.Router) {

		})
	})

	http.ListenAndServe(s.address, r)
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

func (s *Server) createMatch(w http.ResponseWriter, r *http.Request) {

}

/*
func PlayerCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		inGameName := chi.URLParam(r, "inGameName")
		player, err := dbGetPlayer(inGameName)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), "player", player)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getPlayer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	player, ok := ctx.Value("player").(*Player)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	w.Write([]byte(fmt.Sprintf("name:%s", player.Name)))
}
*/
