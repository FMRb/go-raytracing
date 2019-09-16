package main

import (
	"math"
)

// Sphere representation
type Sphere struct {
	center Vec3
	radius float64
}

func (s Sphere) hit(r *Ray, tMin float64, tMax float64, rec *HitRecord) bool {
	oc := r.Origin().Substract(s.center)
	a := r.Direction().Dot(r.Direction())
	b := 2.0 * oc.Dot(r.Direction())
	c := oc.Dot(oc) - (s.radius * s.radius)
	discriminant := b*b - 4*a*c

	if discriminant > 0 {
		temp := (-b - math.Sqrt(discriminant)) / (2 * a)
		if temp < tMax && temp > tMin {
			rec.t = temp
			rec.p = r.PointAtParameter(temp)
			rec.normal = (rec.p.Substract(s.center)).ScalarDivide(s.radius)
			return true
		}
		temp = (-b + math.Sqrt(discriminant)) / (2 * a)
		if temp < tMax && temp > tMin {
			rec.t = temp
			rec.p = r.PointAtParameter(temp)
			rec.normal = (rec.p.Substract(s.center)).ScalarDivide(s.radius)
			return true
		}
	}
	return false
}
