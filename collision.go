package body

import (
	"github.com/polldo/box2d"
)

// Collider is an object capable of listen to collisions against other bodies.
// In order to start listening, a Collider object should be attached to a body.
// This can be done by calling SetObject function of a body.
// If a body has a collider object then the collide function is called every
// time the body collides with another body.
type Collider interface {
	Collide(Body)
}

type CollisionHandler struct {
	pending []collision
}

type collision struct {
	collider Collider
	other    Body
}

func (c *CollisionHandler) BeginContact(contact box2d.B2ContactInterface) {
	a := contact.GetFixtureA().GetBody()
	b := contact.GetFixtureB().GetBody()
	aObj := a.GetUserData()
	bObj := b.GetUserData()

	if aCollider, ok := aObj.(Collider); ok {
		// aCollider.Collide(Body{b})
		c.pending = append(c.pending, collision{aCollider, Body{b}})
	}

	if bCollider, ok := bObj.(Collider); ok {
		// bCollider.Collide(Body{a})
		c.pending = append(c.pending, collision{bCollider, Body{a}})
	}
}

func (c *CollisionHandler) EndContact(contact box2d.B2ContactInterface) {

}

func (c *CollisionHandler) PreSolve(contact box2d.B2ContactInterface, oldManifold box2d.B2Manifold) {
	// possibility to disable the collision

}

func (c *CollisionHandler) PostSolve(contact box2d.B2ContactInterface, impulse *box2d.B2ContactImpulse) {
	// possibility to change effects of collision (friction and restitution)
}

func (c *CollisionHandler) Resolve() {
	for _, coll := range c.pending {
		coll.collider.Collide(coll.other)
	}
	c.pending = nil
}
