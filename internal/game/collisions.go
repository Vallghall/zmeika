package game

import "math/rand"

func (g *Game) CheckFoodCollision() {

	if g.headP == g.foodP {
		for g.foodP = rand.Intn(400) + 1; g.CheckBodyCollision(g.foodP); g.foodP = rand.Intn(400) + 1 {
		}
		g.incrementScore()
		g.growTail()
	}
}

func (g *Game) CheckBodyCollision(p int) bool {
	for _, piece := range g.body {
		if p == piece.pos {
			return true
		}
	}
	return false
}
