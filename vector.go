package physics

import "github.com/polldo/box2d"

type Vec2 struct {
	X, Y float64
}

func V2(x, y float64) Vec2 {
	return Vec2{X: x, Y: y}
}

func (v *Vec2) Set(x, y float64) {
	v.X, v.Y = x, y
}

func (v *Vec2) Add(c float64) *Vec2 {
	v.X, v.Y = v.X+c, v.Y+c
	return v
}

func (v *Vec2) Sub(c float64) *Vec2 {
	v.X, v.Y = v.X-c, v.Y-c
	return v
}

func (v *Vec2) Scale(c float64) *Vec2 {
	v.X, v.Y = v.X*c, v.Y*c
	return v
}

func (v *Vec2) AddV(o Vec2) *Vec2 {
	v.X, v.Y = v.X+o.X, v.Y+o.Y
	return v
}

func (v *Vec2) SubV(o Vec2) *Vec2 {
	v.X, v.Y = v.X-o.X, v.Y-o.Y
	return v
}

func (v *Vec2) ScaleV(o Vec2) *Vec2 {
	v.X, v.Y = v.X*o.X, v.Y*o.Y
	return v
}

func (v Vec2) toB2Vec2() box2d.B2Vec2 {
	return box2d.B2Vec2(v)
}
