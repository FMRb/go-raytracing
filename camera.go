package main

// Camera abstraction
type Camera struct {
	origin          Vec3
	lowerLeftCorner Vec3
	horizontal      Vec3
	vertical        Vec3
}

// NewCamera factory method for Camera type
func NewCamera() Camera {
	return Camera{
		Vec3{0.0, 0.0, 0.0},
		Vec3{-2.0, -1.0, -1.0},
		Vec3{4.0, 0.0, 0.0},
		Vec3{0.0, 2.0, 0.0},
	}
}

// GetRay of the camera
func (c Camera) GetRay(u, v float64) Ray {
	return Ray{
		c.origin,
		c.lowerLeftCorner.Add(c.horizontal.ScalarMultiple(u)).Add(c.vertical.ScalarMultiple(v)).Substract(c.origin),
	}
}
