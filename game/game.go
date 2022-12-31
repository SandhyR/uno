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
	Num      int
	Reverse  bool
	Giliran  *player.Player
	//lmao pisan
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
			"Plus", "Skip", "Reverse",
		}
		switch special[rand.Intn(len(special))] {
		case "Plus":
			pluslist := []int{
				2, 4,
			}
			plus = pluslist[rand.Intn(len(pluslist))]
			if plus == 4 {
				color = "Universal"
			}
			scard = Plus{plus: plus}
			break
		case "Skip":
			scard = Skip{}
			break
		case "Reverse":
			scard = Reverse{}
			break
		}

	}
	id := uuid.NewString()
	card := &Card{color: color, number: number, isSpecial: isspecial, special: scard, Id: id}
	return card, id

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

func (G *Game) CardSession(n int) {
	pdata := G.GetPlayerByNum(n)
	p := pdata.Player
	G.Giliran = p
	p.Message("Sekarang giliran kamu!")

}

func (G *Game) NextPlayer(n int) {
	if !G.Reverse {
		G.Num += n
	} else {
		G.Num -= n
	}
	G.CardSession(G.Num)
}

func (G *Game) HandleSpecialCard(card SpecialCard) {
	switch card.(type) {
	case Plus:
		break
	case Skip:
		G.NextPlayer(2)
	case Reverse:
		if G.Reverse {
			G.Reverse = false
		} else {
			G.Reverse = true
		}
	}
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
	firstcard, _ := RandomizeCard()
	G.LastCard = firstcard
	G.Num = 1
	for _, p := range G.Players {
		for i := 0; i < 5; i++ {
			card, id := RandomizeCard()
			p.Card[id] = card
			p.Player.Message("Kamu mendapat kartu ", card.ToString())
			_, _ = p.Player.Inventory().AddItem(card.ToItem())
		}
		p.Player.Message("Urutan Pemain:\n", message)
		p.Player.Message("Kartu Pertama:", firstcard.ToString())
	}
	G.CardSession(1)
	return true
	//}
	//return false

}

func (G *Game) GetPlayerByNum(n int) *PlayerData {
	count := 1
	for _, p := range G.Players {
		count++
		if count == n {
			return p
		}
	}
	return nil
}
