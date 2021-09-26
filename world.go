package physics

import "github.com/polldo/box2d"

const (
	step               = 1.0 / 60.0
	velocityIterations = 6
	positionIterations = 2
)

var (
	world World
)

func init() {
	gravity := box2d.MakeB2Vec2(0.0, 10)
	world = World{
		box2d.MakeB2World(gravity),
		CollisionHandler{},
	}
	world.SetContactListener(&world.CollisionHandler)
}

type World struct {
	box2d.B2World
	CollisionHandler
}

func Update() {
	world.Step(step, velocityIterations, positionIterations)
	world.Resolve()
}

func SetGravity(x, y float64) {
	world.SetGravity(box2d.B2Vec2{X: x, Y: y})
}

func SetGravityV(g Vec2) {
	world.SetGravity(g.toB2Vec2())
}
