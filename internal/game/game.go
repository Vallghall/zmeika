package game

import (
	"errors"
	"fmt"
	"github.com/Valghall/zmeika/internal/cell"
	c "github.com/Valghall/zmeika/internal/configs"
	"github.com/Valghall/zmeika/internal/food"
	"github.com/Valghall/zmeika/internal/tail"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
)

var (
	headImage *ebiten.Image
	tailImage *ebiten.Image
	foodImage *ebiten.Image
	cellImage *ebiten.Image
	sc        *secretColors

	titleArcadeFont font.Face
	arcadeFont      font.Face
	smallArcadeFont font.Face
	titleTexts      []string
	texts           []string
	secretPhrase    string

	ErrOutOfBorders = errors.New("out of border")
)

func init() {
	foodImage = food.New()
	cellImage = cell.New()
	tailImage = tail.New()
	sc = &secretColors{
		colors: make([]*ebiten.Image, 0),
		len:    0,
	}
	secretPhrase = "Don't push 'C'"

	headImage = ebiten.NewImage(c.CellWidth, c.CellHeight)
	headImage.Fill(color.RGBA{0, 255, 0, 255})

	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	titleArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    c.TitleFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	arcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    c.FontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	smallArcadeFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    c.SmallFontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	headP int
	foodP int

	moveTO int

	direction Direction
	body      []segment
	headPrev  segment

	mode Mode

	bg     *ebiten.Image
	score  int
	secret bool
}

func New() *Game {
	bg := ebiten.NewImage(c.ScreenWidth, c.ScreenHeight)
	bg.Fill(color.RGBA{R: 240, G: 150, B: 100, A: 1})

	return &Game{
		bg:        bg,
		direction: Right,
		mode:      ModeTitle,
		headP:     c.HeadInitialP,
		foodP:     c.FoodInitialP,
		body:      make([]segment, 0),
	}
}

func (g *Game) Update() error {
	if g.mode != ModeGame {
		if g.IsKeyPressed() {
			g.mode = ModeGame
			g.score = 0
		}
		return nil
	}
	g.manageExtraKeys()
	g.ManageControlKey()
	if err := g.Move(); err != nil {
		g.end()
	}

	g.CheckFoodCollision()

	if g.CheckBodyCollision(g.headP) {
		g.end()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	if g.mode == ModeGame {
		g.drawArea(screen)
		g.drawFood(screen)
		g.drawHead(screen)
		g.drawTail(screen)
		ebitenutil.DebugPrintAt(
			screen,
			fmt.Sprintf("SCORE: %d             %s", g.Score(), secretPhrase),
			0,
			c.ScreenHeight,
		)
		return
	}

	screen.DrawImage(g.bg, &ebiten.DrawImageOptions{})
	switch g.mode {
	case ModeTitle:
		titleTexts = []string{"ZMEIKA"}
		texts = []string{"", "", "", "by", "Vallghall", "", "", "", "", "", "PRESS SPACE KEY", "", "OR A/B BUTTON", "", "OR TOUCH SCREEN"}
	case ModeGameOver:
		texts = []string{"", "GAME OVER!", "", "", "SCORE", "", fmt.Sprint(g.Score())}
	}
	for i, l := range titleTexts {
		x := (c.ScreenWidth - len(l)*c.TitleFontSize) / 2
		text.Draw(screen, l, titleArcadeFont, x, (i+4)*c.TitleFontSize, color.RGBA{0, 255, 0, 255})
	}
	for i, l := range texts {
		x := (c.ScreenWidth - len(l)*c.FontSize) / 2
		text.Draw(screen, l, arcadeFont, x, (i+4)*c.FontSize, color.RGBA{0, 255, 0, 255})
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return c.ScreenWidth, int(c.ScreenHeight * 1.05)
}

func (g Game) Score() int {
	return g.score
}

func (g *Game) incrementScore() {
	g.score++
}

func (g *Game) end() {
	bg := ebiten.NewImage(c.ScreenWidth, c.ScreenHeight)
	bg.Fill(color.RGBA{R: 240, G: 150, B: 100, A: 1})
	g.direction = Right
	g.bg = bg
	g.headP = c.HeadInitialP
	g.foodP = c.FoodInitialP
	g.headPrev = segment{}
	g.body = make([]segment, 0)
	sc = &secretColors{
		colors: make([]*ebiten.Image, 0),
		len:    0,
	}
	secretPhrase = "Don't push 'C'"
	titleTexts, texts = []string{}, []string{}
	g.mode = ModeGameOver
}

func (g Game) pos(position int) (float64, float64) {
	x := position % 20
	y := position / 20
	return float64(x) * c.CellWidth, float64(y) * c.CellHeight
}

func (g *Game) TO() {
	g.moveTO += int(X1)
}
