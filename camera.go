package main

import (
	"math"
	"math/rand"
)

// Camera abstraction
type Camera struct {
	origin          Vec3
	lowerLeftCorner Vec3
	horizontal      Vec3
	vertical        Vec3
	u               Vec3
	v               Vec3
	w               Vec3
	lensRadius      float64
}

func randomInUnitDisk() Vec3 {
	var p Vec3
	for {
		p = Vec3{rand.Float64(), rand.Float64(), 0}.ScalarMultiple(2.0).Substract(Vec3{1, 1, 0})
		if p.Dot(p) < 1.0 {
			break
		}
	}
	return p
}

// NewCamera factory method for Camera type
func NewCamera(lookFrom, lookAt, vup Vec3, vfov, aspect, aperture, focusDist float64) Camera {
	lensRadius := aperture / 2
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight
	origin := lookFrom
	w := lookFrom.Substract(lookAt).UnitVector()
	u := vup.Cross(w).UnitVector()
	v := w.Cross(u)
	lowerLeftCorner := origin.Substract(
		u.ScalarMultiple(halfWidth * focusDist)).
		Substract(v.ScalarMultiple(halfHeight * focusDist)).
		Substract(w.ScalarMultiple(focusDist))
	horizontal := u.ScalarMultiple(2 * halfWidth * focusDist)
	vertical := v.ScalarMultiple(2 * halfHeight * focusDist)
	return Camera{origin, lowerLeftCorner, horizontal, vertical, u, v, w, lensRadius}
}

// GetRay of the camera
func (c Camera) GetRay(u, v float64) Ray {
	rd := randomInUnitDisk().ScalarMultiple(c.lensRadius)
	offset := c.u.ScalarMultiple(rd.X).Add(c.v.ScalarMultiple(rd.Y))
	return Ray{
		c.origin.Add(offset),
		c.lowerLeftCorner.Add(c.horizontal.ScalarMultiple(u)).
			Add(c.vertical.ScalarMultiple(v)).Substract(c.origin).Substract(offset),
	}
}
