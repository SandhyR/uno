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
			pgame.StartGame()

		}
	}
}
