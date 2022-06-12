package tail

import (
	"github.com/Valghall/zmeika/internal/configs"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	_ "image/png"
)

type TailSegment *ebiten.Image

func New() TailSegment {
	//f, _, err := ebitenutil.NewImageFromFile(configs.TailPath)
	f := ebiten.NewImage(configs.CellWidth, configs.CellHeight)
	/*if err != nil {
		log.Fatalln(err)
	}

	*/
	f.Fill(color.RGBA{
		R: 50,
		G: 100,
		B: 50,
		A: 255,
	})
	return f
}
