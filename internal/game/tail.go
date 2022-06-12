package game

import (
	c "github.com/Valghall/zmeika/internal/configs"
)

type segment struct {
	direction Direction
	pos       int
}

func (g *Game) growTail() {
	var last segment
	if len(g.body) == 0 {
		g.body = append(g.body, segment{
			direction: g.direction,
			pos: g.findPosition(segment{
				direction: g.direction,
				pos:       g.headP,
			}),
		})
		return
	}
	last = g.body[len(g.body)-1]
	g.body = append(g.body, segment{
		direction: last.direction,
		pos:       g.findPosition(last),
	})
	sc.add()
}

func (g *Game) findPosition(seg segment) int {
	switch seg.direction {
	case Right:
		if clouse := (seg.pos - 1) % c.CiaR; (clouse != c.CiaR-1 && clouse > -1) && seg.pos-1 != g.headP {
			if !g.CheckBodyCollision(seg.pos - 1) {
				return seg.pos - 1
			}
		}
		fallthrough
	case Down:
		if (seg.pos-c.CiaR >= 0) && seg.pos-c.CiaR != g.headP {
			if !g.CheckBodyCollision(seg.pos - c.CiaR) {
				return seg.pos - c.CiaR
			}
		}
		fallthrough
	case Left:
		if ((seg.pos+1)%c.CiaR != 0) && seg.pos+1 != g.headP {
			if !g.CheckBodyCollision(seg.pos + 1) {
				return seg.pos + 1
			}
		}
		fallthrough
	case Up:
		if (seg.pos+c.CiaR <= c.AoC) && seg.pos+c.CiaR != g.headP {
			if !g.CheckBodyCollision(seg.pos + c.CiaR) {
				return seg.pos + c.CiaR
			}
		}
	}
	return seg.pos + 1
}
