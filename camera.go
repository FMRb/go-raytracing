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
func NewCamera(lookFrom, lookAt, vup Vec3, vfov, aspect float64) Camera {
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight
	origin := lookFrom
	w := lookFrom.Substract(lookAt).UnitVector()
	u := vup.Cross(w).UnitVector()
	v := w.Cross(u)
	lowerLeftCorner := Vec3{-halfWidth, -halfHeight, -1.0}
	lowerLeftCorner = origin.Substract(u.ScalarMultiple(halfWidth)).Substract(v.ScalarMultiple(halfHeight)).Substract(w)
	horizontal := u.ScalarMultiple(2 * halfWidth)
	vertical := v.ScalarMultiple(2 * halfHeight)
	return Camera{origin, lowerLeftCorner, horizontal, vertical}
}

// GetRay of the camera
func (c Camera) GetRay(u, v float64) Ray {
	return Ray{
		c.origin,
		c.lowerLeftCorner.Add(c.horizontal.ScalarMultiple(u)).Add(c.vertical.ScalarMultiple(v)).Substract(c.origin),
	}
}
