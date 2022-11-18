package game

import "fmt"

type Card struct {
	color     string
	number    int
	isspecial bool
	special   SpecialCard
}

func (C *Card) ToString() string {
	if C.isspecial {
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
		return fmt.Sprintf("%v Warna %v", text, C.color)
	}
	return fmt.Sprintf("Nomor %v Warna %v", C.number, C.color)
}
