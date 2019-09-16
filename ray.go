package main

// Ray define by two point A(Origin) and B(Direction)
type Ray struct {
	A Vec3
	B Vec3
}

// Origin point of the ray in a 3D space
func (r *Ray) Origin() Vec3 {
	return r.A
}

// Direction point of the ray in a 3D space
func (r *Ray) Direction() Vec3 {
	return r.B
}

// PointAtParameter calculate point along the ray using a real number
func (r *Ray) PointAtParameter(t float64) Vec3 {
	return r.A.Add(r.B.ScalarMultiple(t))
}
