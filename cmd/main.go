package main

import (
	"context"
	"fmt"

	"github.com/erupshis/golang-integration-developer-test/internal/service/client"
)

func main() {

	defClient := client.NewDefault("")

	games, _ := defClient.GetGames(context.Background(), "pc")
	fmt.Printf("%+v\n", games)

}
