package game

import (
	"github.com/Valghall/zmeika/internal/configs"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type Mode int

const (
	ModeTitle Mode = iota
	ModeGame
	ModeGameOver
)

type Game struct {
	HeadPositionX float64
	HeadPositionY float64

	bg *ebiten.Image
}

func New() *Game {
	bg := ebiten.NewImage(configs.ScreenWidth, configs.ScreenHeight)
	bg.Fill(color.RGBA{R: 240, G: 150, B: 100, A: 1})

	return &Game{
		float64(configs.ScreenWidth)/2 - 20,
		float64(configs.ScreenHeight) / 2,
		bg,
	}
}

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.bg, &ebiten.DrawImageOptions{})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return configs.ScreenWidth, configs.ScreenHeight
}
