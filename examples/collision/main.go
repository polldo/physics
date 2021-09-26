package main

import (
	"image/color"

	"physics"

	"github.com/go-p5/p5"
)

const (
	width  = 640
	height = 640

	rad = 5.
)

var (
	ball1 Ball
	ball2 Ball

	wall   Wall
	wallUp Wall
)

type Ball struct {
	physics.Body
	coll bool
}

func (b *Ball) Collide(other physics.Body) {
	b.Impulse(physics.V2(0, -30))
}

func (b *Ball) Update() {
}

type Wall struct {
	physics.Body
}

func setup() {
	p5.Canvas(width, height)
	p5.Stroke(nil)
	p5.Fill(color.RGBA{R: 255, A: 204})

	physics.SetGravity(0.3, 30)

	ball1.Body = physics.NewBody(physics.Type(physics.DynamicBody), physics.Circle(rad))
	ball1.SetPosition(100, 100)
	ball1.SetObject(&ball1)
	ball1.SetCollisionFilter(0x2, 0xffff)

	ball2.Body = physics.NewBody(physics.Type(physics.DynamicBody), physics.Circle(rad))
	ball2.SetPosition(200, 100)

	wall.Body = physics.NewBody(physics.Type(physics.StaticBody), physics.Rectangle(width, 20))
	wall.SetPosition(0, 140)

	wallUp.Body = physics.NewBody(physics.Type(physics.StaticBody), physics.Rectangle(width, 20))
	wallUp.SetCollisionFilter(0x2, 0x2)
	wallUp.SetPosition(0, 120)
}

func main() {
	p5.Run(setup, draw)
}

func draw() {
	physics.Update()
	ball1.Update()
	ball2.Update()

	p5.Background(color.Gray{Y: 220})

	p5.Fill(color.RGBA{R: 255, A: 204})
	p5.Ellipse(ball1.X(), ball1.Y(), 2*rad, 2*rad)
	p5.Ellipse(ball2.X(), ball2.Y(), 2*rad, 2*rad)

	// fmt.Println(wall.X(), wall.Y(), wall.X()+width, wall.Y())
	p5.StrokeWidth(10)
	p5.Stroke(color.RGBA{R: 255, A: 204})
	p5.Line(wall.X(), wall.Y(), wall.X()+width, wall.Y())
	p5.Line(wallUp.X(), wallUp.Y(), wallUp.X()+width, wallUp.Y())
}
