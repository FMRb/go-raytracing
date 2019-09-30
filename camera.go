package main

import "math"

// Camera abstraction
type Camera struct {
	origin          Vec3
	lowerLeftCorner Vec3
	horizontal      Vec3
	vertical        Vec3
}

// NewCamera factory method for Camera type
func NewCamera(vfov, aspect float64) Camera {
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight
	lowerLeftCorner := Vec3{-halfWidth, -halfHeight, -1.0}
	horizontal := Vec3{2 * halfWidth, 0.0, 0.0}
	vertical := Vec3{0.0, 2 * halfHeight, 0.0}
	origin := Vec3{0.0, 0.0, 0.0}
	return Camera{origin, lowerLeftCorner, horizontal, vertical}
}

// GetRay of the camera
func (c Camera) GetRay(u, v float64) Ray {
	return Ray{
		c.origin,
		c.lowerLeftCorner.Add(c.horizontal.ScalarMultiple(u)).Add(c.vertical.ScalarMultiple(v)).Substract(c.origin),
	}
}
