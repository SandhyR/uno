package game

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
	session := GetSession(p.player.Name())
	item, _ := p.player.HeldItems()
	uuid, ok := item.Value("uuid")

	if !ok {
		return
	}
	special, ok := item.Value("special")
	if session.game.Giliran.Name() == p.player.Name() {
		delete(session.game.Players[p.player.Name()].Card, uuid.(string))
		_ = p.player.Inventory().RemoveItem(item)
		session.game.NextPlayer(1)
		if ok {
			session.game.HandleSpecialCard(special.(SpecialCard))
			return
		}
	} else {
		p.player.Message("Ini bukan giliran kamu!")
	}
}
