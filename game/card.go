package game

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/item"
)

type Card struct {
	color     string
	number    int
	isSpecial bool
	special   SpecialCard
}

func (C *Card) ToString() string {
	if C.isSpecial {
		var text string
		switch C.special.(type) {
		case *Reverse:
			text = "Reverse"
			break
		case *Skip:
			text = "Skip"
			break
		case *Plus:
			text = "Plus"
			break
		}
		return fmt.Sprintf("%v %v", text, C.color)
	}
	return fmt.Sprintf("#%v %v", C.number, C.color)
}

func (C *Card) ToItem() item.Stack {
	if C.isSpecial {
		return item.NewStack(item.Paper{}, 1).WithCustomName(C.ToString()).WithValue("color", C.color).WithValue("special", C.special)
	}
	return item.NewStack(item.Paper{}, 1).WithCustomName(C.ToString()).WithValue("special", false).WithValue("color", C.color).WithValue("number", C.number)
}
