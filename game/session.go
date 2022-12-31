package game

import "github.com/df-mc/dragonfly/server/player"

var sessions map[string]*Session

type Session struct {
	game *Game
}

func CreateSession(p *player.Player) {
	sessions[p.Name()] = &Session{}
}

func GetSession(p string) *Session {
	return sessions[p]

}

func RemoveSession(p string) {
	delete(sessions, p)
}
