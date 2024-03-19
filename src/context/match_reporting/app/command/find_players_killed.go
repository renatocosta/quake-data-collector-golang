package command

import (
	"context"
	"fmt"

	"github.com/ddd/crosscutting/building_blocks/app"
	players_killed "github.com/ddd/src/context/match_reporting/domain/model/player_killed"
)

type FindPlayersKilledCommand struct {
	Data [][]string
}

type FindPlayersKilledHandler app.CommandHandler[FindPlayersKilledCommand, []string]

type findPlayersKilledHandler struct {
}

func NewFindPlayersKilledHandler() app.CommandHandler[FindPlayersKilledCommand, []string] {
	return findPlayersKilledHandler{}
}

func (h findPlayersKilledHandler) Handle(ctx context.Context, cmd FindPlayersKilledCommand) ([]string, error) {

	playersKilled := players_killed.NewPlayerKilled(players_killed.NewKilledPlayer())

	for _, row := range cmd.Data {
		playersKilled.Compute(players_killed.NewMatcher(row[0], row[1]))
	}

	jsonData := `{
		"game_1": {
			"total_kills": 45,
			"players": ["Dono da bola", "Isgalamido", "Zeh"],
			"kills": {
				"Dono da bola": 5,
				"Isgalamido": 18,
				"Zeh": 20
			}
		}
	}`

	fmt.Print(jsonData)

	return []string{}, nil
}
