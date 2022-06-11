package cell

import (
	"github.com/Valghall/zmeika/internal/configs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
)

type Cell *ebiten.Image

func New() Cell {
	f, _, err := ebitenutil.NewImageFromFile(configs.CellPath)

	if err != nil {
		log.Fatalln(err)
	}
	return ebiten.NewImageFromImage(f)
}
