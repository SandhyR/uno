package handler

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
)

type PlayerHandler struct {
	player.NopHandler
	player *player.Player
}

func AddToHandler(player *player.Player) *PlayerHandler {
	handler := &PlayerHandler{
		player: player,
	}
	return handler
}

func (p *PlayerHandler) HandleItemUse(context *event.Context) {
	//TODO: Handle Card
}
