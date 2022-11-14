package cmd

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"uno/game"
)

type CreateGame struct {
}

// Run will be called when the player runs the command. In this case, we will print the number back to the player
func (c CreateGame) Run(source cmd.Source, output *cmd.Output) {
	if source, ok := source.(*player.Player); ok {
		id := game.CreateNewGame(source)
		output.Printf("Membuat game baru id: %v", id)
	}

}
