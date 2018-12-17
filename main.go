package main

import (
	"github.com/bavaz1/go-pubg/server"
)

func main() {
	/* ctx := context.Background()
	playerName := "bAVAZ1"
	client := http.Client{}
	resp, err := sdk.GetPlayer(ctx, playerName, &client)
	if err != nil {
		panic(err)
	}

	var bavaz1LastMatch string = resp.Data[0].Relationships.Matches.Data[0].ID

	resp2, err := sdk.GetMatch(ctx, bavaz1LastMatch, &client)
	if err != nil {
		panic(err)
	} */

	server.ListenAndServe(8080)
}
