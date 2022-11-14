package game

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/google/uuid"
	"math/rand"
	"time"
	"uno/utils"
)

var Games = map[string]*Game{}

type Game struct {
	Creator  *player.Player
	Players  map[string]*PlayerData
	Start    bool
	Id       string
	LastCard *Card
}

type PlayerData struct {
	Player *player.Player
	Card   map[int]*Card
}

type Card struct {
	color     string
	number    int
	isspecial bool
	special   SpecialCard
}

func CreateNewGame(creator *player.Player) string {
	id := uuid.New().String()
	Games[id] = &Game{Start: false, Id: id, Creator: creator, Players: map[string]*PlayerData{creator.Name(): &PlayerData{creator, map[int]*Card{}}}}
	return id
}

func JoinGame(player *player.Player, id string) bool {
	if _, ok := Games[id]; ok {
		Games[id].Players[player.Name()] = &PlayerData{Player: player, Card: map[int]*Card{}}
		return true
	}
	return false
}

func RandomizeCard() *Card {
	rand.Seed(time.Now().Unix())
	colorlist := []string{
		"Merah",
		"Kuning",
		"Hijau",
		"Biru",
	}
	n := rand.Int() % len(colorlist)
	color := colorlist[n]
	number := utils.RandomNumber()
	isspecial := utils.RandomBool()
	plus := 0
	var scard SpecialCard
	if isspecial {
		number = 0
		rand.Seed(time.Now().Unix())
		special := []string{
			"Plus", "Skip",
		}
		n = rand.Int() % len(special)
		if special[n] == "Plus" {
			rand.Seed(time.Now().Unix())
			pluslist := []int{
				2, 4,
			}
			n = rand.Int() % len(pluslist)
			plus = pluslist[n]
			if plus == 4 {
				color = "Universal"
			}
			scard = &Plus{plus: plus}
		} else {
			scard = &Skip{}
		}

	}
	card := &Card{color: color, number: number, isspecial: isspecial, special: scard}
	return card

}

func GetGame(p *player.Player) (*Game, bool) {
	for _, game := range Games {
		for _, playerdata := range game.Players {
			if playerdata.Player.Name() == p.Name() {
				return game, true
			}
		}

	}
	return nil, false
}

func (G *Game) StartGame() {
	G.Start = true

}
