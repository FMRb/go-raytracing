package main

import (
	"math"
)

// Reflect calculation of the reflection vector
// Parameter with incident vector v and normal vector n
func Reflect(v Vec3, n Vec3) Vec3 {
	return v.Substract(n.ScalarMultiple(2.0 * v.Dot(n)))
}

// Refract calculate the refreaction vector from
// an incident vector v with the normal vector n and the
// refraction incident of the two materials, incident_material/target_material
func Refract(v Vec3, n Vec3, niOverNt float64, refracted *Vec3) bool {
	uv := v.UnitVector()
	dt := uv.Dot(n)
	discriminant := 1.0 - niOverNt*niOverNt*(1.0-dt*dt)
	if discriminant > 0 {
		param1 := uv.Substract(n.ScalarMultiple(dt)).ScalarMultiple(niOverNt)
		param2 := n.ScalarMultiple(math.Sqrt(discriminant))
		*refracted = param1.Substract(param2)
		return true
	}
	return false
}

// Schlick calculate reflectivity depending on the angle
func Schlick(cosine float64, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
