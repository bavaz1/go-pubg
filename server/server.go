package server

import (
	"fmt"
	"net/http"

	"github.com/bavaz1/go-pubg/database"
	"github.com/go-chi/chi"
)

type Server struct {
	address string
	storage database.Storage
}

func (s *Server) Listen() {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/players", func(r chi.Router) {
			r.Post("/{name}", s.createPlayer)
		})
		r.Route("/match", func(r chi.Router) {

		})
	})

	http.ListenAndServe(s.address, r)
}

func (s *Server) createPlayer(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name") // in game name

	// ctx := r.Context()

	w.Write([]byte(fmt.Sprintf("name:%s", name)))
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
