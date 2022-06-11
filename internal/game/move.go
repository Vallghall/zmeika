package game

import (
	c "github.com/Valghall/zmeika/internal/configs"
	"log"
)

func (g *Game) Move() error {
	if g.moveTO != 0 {
		g.moveTO--
		return nil
	}
	g.TO()
	g.headPrev = segment{
		direction: g.direction,
		pos:       g.headP,
	}

	defer g.moveTail()

	switch g.direction {
	case Left:
		g.headP--
		if clouse := g.headP % c.CiaR; clouse == c.CiaR-1 || clouse == -1 {
			log.Printf("Snake died at cell No %d, moving in direction %d\n", g.headP, g.direction)
			return ErrOutOfBorders
		}
	case Up:
		g.headP -= c.CiaR
		if g.headP < 0 {
			log.Printf("Snake died at cell No %d, moving in direction %d\n", g.headP, g.direction)
			return ErrOutOfBorders
		}
	case Right:
		g.headP++
		if g.headP%c.CiaR == 0 {
			log.Printf("Snake died at cell No %d, moving in direction %d\n", g.headP, g.direction)
			return ErrOutOfBorders
		}
	case Down:
		g.headP += c.CiaR
		if g.headP > c.AoC {
			log.Printf("Snake died at cell No %d, moving in direction %d\n", g.headP, g.direction)
			return ErrOutOfBorders
		}
	}

	return nil
}

func (g *Game) moveTail() {
	l := len(g.body)
	if l == 0 {
		return
	}
	temp := make([]segment, l)
	copy(temp, g.body)
	for i := l - 1; i > 0; i-- {
		g.body[i] = temp[i-1]
	}
	g.body[0] = g.headPrev
}
