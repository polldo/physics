package body

import "github.com/polldo/box2d"

const (
	step               = 1.0 / 60.0
	velocityIterations = 6
	positionIterations = 2
)

var (
	world World
)

func V2(x, y float64) box2d.B2Vec2 {
	return box2d.B2Vec2{X: x, Y: y}
}

func init() {
	gravity := box2d.MakeB2Vec2(0.0, 10)
	world = World{box2d.MakeB2World(gravity)}
}

type World struct {
	box2d.B2World
}

func Update() {
	world.Step(step, velocityIterations, positionIterations)
}

func SetGravity(g box2d.B2Vec2) {
	world.SetGravity(g)
}
