package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bavaz1/go-pubg/sdk"
)

func main() {
	ctx := context.Background()
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
	}

	fmt.Println(resp2)
}
