package body

import (
	"github.com/polldo/box2d"
)

// Collider is an entity interested in collision callbacks for a
// certain body. It can be enabled by calling EnableBodyCollider.
type Collider interface {
	Collide(interface{})
}

// BodyCollider is for entities that embed a body and are interested
// in their collision callback.
type BodyCollider interface {
	Collider
	// We use an unexported function to be sure this is a body
	setUserData(interface{})
}

// RegisterBodyCollider registers an entity which wants to receive callbacks
// when its body collides with other bodies.
func RegisterBodyCollider(collider BodyCollider) {
	collider.setUserData(collider)
}

// RegisterCollider registers an entity which wants to receive callbacks
// when the passed body collides with other bodies.
func RegisterCollider(collider Collider, body Body) {
	body.setUserData(collider)
}

type CollisionHandler struct{}

func (c *CollisionHandler) BeginContact(contact box2d.B2ContactInterface) {
	a := contact.GetFixtureA().GetBody().GetUserData()
	b := contact.GetFixtureB().GetBody().GetUserData()

	if aCollider, ok := a.(Collider); ok {
		aCollider.Collide(b)
	}

	if bCollider, ok := b.(Collider); ok {
		bCollider.Collide(b)
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
