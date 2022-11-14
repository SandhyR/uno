package game

type SpecialCard interface {
}

type Plus struct {
	SpecialCard
	plus int
}

type Skip struct {
	SpecialCard
}
