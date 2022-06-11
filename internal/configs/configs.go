package configs

const (
	Title         = "Zmeika by Vallghall"
	ScreenWidth   = 640
	ScreenHeight  = 640
	TitleFontSize = FontSize * 1.5
	FontSize      = 24
	SmallFontSize = FontSize / 2

	LogoPath = "./images/logo2.png"
	FoodPath = "./images/apple.png"
	CellPath = "./images/cell.png"
	TailPath = "./images/tail.png"

	HeadInitialP = 205
	FoodInitialP = 215

	CiaR = 20          // Cells in a Row
	CiaC = 20          // Cells in a Column
	AoC  = CiaR * CiaC // Amount of Cells

	CellWidth  = ScreenWidth / CiaR
	CellHeight = ScreenHeight / CiaC
)
