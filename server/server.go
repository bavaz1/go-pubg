package server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
)

func ListenAndServe(port int) {
	r := chi.NewRouter()
	p := strconv.Itoa(port)
	var sb strings.Builder
	sb.WriteString(":")
	sb.WriteString(p)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Route("/players", func(r chi.Router) {
		r.With(paginate).Get("/", listPlayers) // GET /players

		r.Post("/", createPlayer) // POST /players

		// Subrouters:
		r.Route("/{inGameName}", func(r chi.Router) {
			r.Use(PlayerCtx)
			r.Get("/", getPlayer)
			r.Put("/", updatePlayer)
			r.Delete("/", deletePlayer)
		})
	})

	http.ListenAndServe(sb.String(), r)
}

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
