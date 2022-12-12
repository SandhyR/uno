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
				ok = pgame.StartGame()
				if ok {
					output.Printf("Berhasil memulai game")
				} else {
					output.Printf("Pemain tidak cukup!")
				}
			} else {
				output.Printf("Kamu bukan creator game ini!")
			}
		} else {
			output.Printf("Kamu tidak masuk game manapun")
		}
	}
}
