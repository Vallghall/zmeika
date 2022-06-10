package main

import (
	"github.com/Valghall/zmeika/internal/configs"
	"github.com/Valghall/zmeika/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image"
	_ "image/png"
	"log"
)

func init() {
	logo, _, err := ebitenutil.NewImageFromFile(configs.LogoPath)
	if err != nil {
		log.Fatalln(err)
	}

	ebiten.SetWindowSize(configs.ScreenWidth, int(configs.ScreenHeight*1.05))
	ebiten.SetWindowTitle(configs.Title)
	ebiten.SetWindowIcon([]image.Image{logo})
}

func main() {

	log.Fatalln(ebiten.RunGame(game.New()))
}
