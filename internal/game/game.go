package game

import (
	"errors"
	"fmt"
	"github.com/Valghall/zmeika/internal/cell"
	c "github.com/Valghall/zmeika/internal/configs"
	"github.com/Valghall/zmeika/internal/food"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
	"math/rand"
)

type Mode int

const (
	ModeTitle Mode = iota
	ModeGame
	ModeGameOver
)

type Direction int

const (
	Left Direction = iota
	Up
	Right
	Down
)

type SpeedLimiter int

const (
	X1 SpeedLimiter = 60/10 + iota
	X2
	X3
)

var (
	headImage *ebiten.Image
	foodImage *ebiten.Image
	cellImage *ebiten.Image

	titleArcadeFont font.Face
	arcadeFont      font.Face
	smallArcadeFont font.Face
	titleTexts      []string
	texts           []string

	ErrOutOfBorders = errors.New("out of border")
)

func init() {
	foodImage = food.New()
	cellImage = cell.New()

	headImage = ebiten.NewImage(c.ScreenWidth/20, c.ScreenHeight/20)
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

	mode Mode

	bg    *ebiten.Image
	score int
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
	}
}

func (g *Game) Update() error {
	if g.mode != ModeGame {
		if g.IsKeyPressed() {
			g.mode = ModeGame
		}
		return nil
	}
	g.ManageControlKey()
	if err := g.Move(); err != nil {
		g.end()
	}

	g.CheckFoodCollision()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	if g.mode == ModeGame {
		g.drawArea(screen)
		g.drawHead(screen)
		g.drawFood(screen)
		ebitenutil.DebugPrintAt(
			screen,
			fmt.Sprintf("SCORE: %d", g.Score()),
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
		texts = []string{"", "GAME OVER!"}
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

func (g *Game) drawArea(screen *ebiten.Image) {

	for i := 0; i < c.AoC; i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(g.pos(i))
		screen.DrawImage(cellImage, op)
	}
}

func (g Game) Score() int {
	return g.score
}

func (g *Game) incrementScore() {
	g.score++
}

func (g Game) IsKeyPressed() bool {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) ||
		inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		return true
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) ||
		inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return true
	}

	return false
}

func (g *Game) ManageControlKey() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) ||
		inpututil.IsKeyJustPressed(ebiten.KeyA) {
		if g.direction != Right {
			g.direction = Left
		}

		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) ||
		inpututil.IsKeyJustPressed(ebiten.KeyW) {
		if g.direction != Down {
			g.direction = Up
		}

		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) ||
		inpututil.IsKeyJustPressed(ebiten.KeyD) {
		if g.direction != Left {
			g.direction = Right
		}

		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) ||
		inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if g.direction != Up {
			g.direction = Down
		}

		return
	}
}

func (g *Game) Move() error {
	if g.moveTO != 0 {
		g.moveTO--
		return nil
	}
	g.TO()

	switch g.direction {
	case Left:
		g.headP--
		if clouse := g.headP % c.CiaR; clouse == c.CiaR-1 || clouse == -1 {
			log.Printf("Snake died at cell No %d, moving in direction %d\n", g.headP, g.direction)
			return ErrOutOfBorders
		}
	case Up:
		g.headP -= c.CiaR
		if g.headP < 1 {
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

func (g *Game) CheckFoodCollision() {

	if g.headP == g.foodP {
		g.foodP = rand.Intn(40) + 1
		g.incrementScore()
	}
}

func (g *Game) end() {
	bg := ebiten.NewImage(c.ScreenWidth, c.ScreenHeight)
	bg.Fill(color.RGBA{R: 240, G: 150, B: 100, A: 1})
	g.score = 0
	g.direction = Right
	g.bg = bg
	g.headP = c.HeadInitialP
	g.foodP = c.FoodInitialP
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
