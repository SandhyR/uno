package game

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/google/uuid"
	"math/rand"
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
	Card   map[string]*Card
}

func CreateNewGame(creator *player.Player) string {
	id := uuid.New().String()
	Games[id] = &Game{Start: false, Id: id, Creator: creator, Players: map[string]*PlayerData{creator.Name(): &PlayerData{creator, map[string]*Card{}}}}
	return id
}

func JoinGame(player *player.Player, id string) bool {
	if _, ok := Games[id]; ok {
		Games[id].Players[player.Name()] = &PlayerData{Player: player, Card: map[string]*Card{}}
		return true
	}
	return false
}

func RandomizeCard() (*Card, string) {
	colorlist := []string{
		"Merah",
		"Kuning",
		"Hijau",
		"Biru",
	}
	color := colorlist[rand.Intn(len(colorlist))]
	number := utils.RandomNumber()
	isspecial := utils.RandomBool()
	plus := 0
	var scard SpecialCard
	if isspecial {
		number = 0
		special := []string{
			"Plus", "Skip",
		}
		if special[rand.Intn(len(special))] == "Plus" {
			pluslist := []int{
				2, 4,
			}
			plus = pluslist[rand.Intn(len(pluslist))]
			if plus == 4 {
				color = "Universal"
			}
			scard = &Plus{plus: plus}
		} else {
			scard = &Skip{}
		}

	}
	card := &Card{color: color, number: number, isSpecial: isspecial, special: scard}
	return card, uuid.NewString()

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

func (G *Game) StartGame() bool {
	G.Start = true
	//if len(G.Players) > 1 {
	var message string
	count := 1
	for _, p := range G.Players {
		message += fmt.Sprintf("%v. %v \n", count, p.Player.Name())
		count++
	}

	for _, p := range G.Players {
		for i := 0; i < 5; i++ {
			card, id := RandomizeCard()
			fmt.Println(card)
			p.Card[id] = card
			p.Player.Message("Kamu mendapat kartu ", card.ToString())
			_, _ = p.Player.Inventory().AddItem(card.ToItem())
		}
		p.Player.Message("Urutan Pemain:\n", message)
	}
	return true
	//}
	//return false

}
