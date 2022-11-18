package cmd

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"uno/game"
)

type StartCommand struct {
}

func (c StartCommand) Run(source cmd.Source, output *cmd.Output) {
	if source, ok := source.(*player.Player); ok {
		pgame, ok := game.GetGame(source)
		if ok {
			if pgame.Creator.Name() == source.Name() {
				pgame.StartGame()
			} else {
				output.Printf("Kamu bukan creator game ini!")
			}
		} else {
			output.Printf("Kamu tidak masuk game manapun")
		}
	}
}
