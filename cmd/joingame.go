package cmd

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"uno/game"
)

type JoinGame struct {
	Id string `cmd:"id"`
}

func (c JoinGame) Run(source cmd.Source, output *cmd.Output) {
	if source, ok := source.(*player.Player); ok {
		ok = game.JoinGame(source, c.Id)
		if ok {
			output.Printf("Berhasil join!")
		} else {
			output.Printf("Game dengan id: %v tidak ditemukan", c.Id)
		}
	}
}
