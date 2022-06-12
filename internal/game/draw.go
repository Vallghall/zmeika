package game

import (
	c "github.com/Valghall/zmeika/internal/configs"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math/rand"
)

func (g *Game) drawArea(screen *ebiten.Image) {

	for i := 0; i < c.AoC; i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(g.pos(i))
		screen.DrawImage(cellImage, op)
	}
}

func (g *Game) drawHead(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{Filter: ebiten.FilterLinear}
	op.GeoM.Translate(g.pos(g.headP))
	screen.DrawImage(headImage, op)
}

func (g Game) drawFood(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{Filter: ebiten.FilterLinear}
	op.GeoM.Translate(g.pos(g.foodP))
	screen.DrawImage(foodImage, op)
}

type secretColors struct {
	colors []*ebiten.Image
	len    int
}

func (sc *secretColors) add() {
	sc.len++
	sc.colors = append(sc.colors, tailImage)
}

func (g *Game) drawTail(screen *ebiten.Image) {
	if len(g.body) == 0 {
		return
	}

	if g.secret {

		for len(g.body) > sc.len {
			sc.add()
		}

		if g.moveTO == 0 {
			for i := 0; i < sc.len; i++ {
				sc.colors[i] = ebiten.NewImage(c.CellWidth, c.CellHeight)
				sc.colors[i].Fill(color.RGBA{
					R: uint8(rand.Intn(256)),
					G: uint8(rand.Intn(256)),
					B: uint8(rand.Intn(256)),
					A: 255,
				})
			}
		}
	}

	for i, piece := range g.body {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(g.pos(piece.pos))

		if g.secret {
			screen.DrawImage(sc.colors[i], op)
			continue
		}

		screen.DrawImage(tailImage, op)
	}
}
