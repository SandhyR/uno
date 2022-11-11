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
	creator  *player.Player
	players  map[string]*PlayerData
	start    bool
	id       string
	lastcard *Card
}

type PlayerData struct {
	player *player.Player
	card   map[int]*Card
}

type Card struct {
	color  string
	number int
	isplus bool
	plus   int
}

func CreateNewGame(creator *player.Player) {
	id := uuid.New().String()
	Games[id] = &Game{start: false, id: id, creator: creator, players: map[string]*PlayerData{creator.Name(): &PlayerData{creator, map[int]*Card{}}}}
}

func JoinGame(player *player.Player, id string) bool {
	if _, ok := Games[id]; ok {
		Games[id].players[player.Name()] = &PlayerData{player: player, card: map[int]*Card{}}
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
	isplus := utils.RandomBool()
	plus := 0
	if isplus {
		number = 0
		rand.Seed(time.Now().Unix())
		pluslist := []int{
			2, 4,
		}
		n = rand.Int() % len(pluslist)
		plus = pluslist[n]
		if plus == 4 {
			color = "Universal"
		}
	}
	card := &Card{color: color, number: number, isplus: isplus, plus: plus}
	return card

}
