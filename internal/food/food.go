package food

import (
	"github.com/Valghall/zmeika/internal/configs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"log"
)

type Food *ebiten.Image

func New() Food {
	f, _, err := ebitenutil.NewImageFromFile(configs.FoodPath)

	if err != nil {
		log.Fatalln(err)
	}
	return ebiten.NewImageFromImage(f)
}
