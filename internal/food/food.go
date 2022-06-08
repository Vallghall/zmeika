package food

import (
	"github.com/Valghall/zmeika/internal/configs"
	"github.com/hajimehoshi/ebiten/v2"
	"image/png"
	"log"
	"os"
)

type Food struct {
	sprite *ebiten.Image
}

func New() *Food {
	f, _ := os.Open(configs.FoodPath)
	defer f.Close()

	sprite, err := png.Decode(f)
	if err != nil {
		log.Fatalln(err)
	}
	return &Food{ebiten.NewImageFromImage(sprite)}
}
