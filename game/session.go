package game

import "github.com/df-mc/dragonfly/server/player"

var sessions map[string]*Session

type Session struct {
	game *Game
}

func CreateSession(p *player.Player) {
	sessions[p.Name()] = &Session{}
}

func RemoveSession(p string) {
	delete(sessions, p)
}
