package game

type Mode int

const (
	ModeTitle Mode = iota
	ModeGame
	ModeGameOver
)

type Direction int

const (
	Left Direction = iota
	Up
	Right
	Down
)

type SpeedLimiter int

const (
	X1 SpeedLimiter = 60/10 + iota
	X2
	X3
)
