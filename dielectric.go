package main

import "math/rand"

// Dielectric material, calculating reflection or refraction
type Dielectric struct {
	refIdx float64
}

func (d Dielectric) scatter(rIn *Ray, rec *HitRecord, attenuation *Vec3, scattered *Ray) bool {
	var outwardNormal Vec3
	reflected := Reflect(rIn.Direction(), rec.normal)
	var niOverNt float64
	*attenuation = Vec3{1.0, 1.0, 1.0}
	var refracted Vec3
	var reflectedProb float64
	var cosine float64
	// Use of refraction air 1.0 for the space outside of the hitting object
	if rIn.Direction().Dot(rec.normal) > 0 {
		outwardNormal = rec.normal.ScalarMultiple(-1)
		niOverNt = d.refIdx
		cosine = d.refIdx * rIn.Direction().Dot(rec.normal) / rIn.Direction().Length()
	} else {
		outwardNormal = rec.normal
		niOverNt = 1.0 / d.refIdx
		cosine = -(rIn.Direction().Dot(rec.normal) / rIn.Direction().Length())
	}

	if Refract(rIn.Direction(), outwardNormal, niOverNt, &refracted) {
		reflectedProb = Schlick(cosine, d.refIdx)
	} else {
		*scattered = Ray{rec.p, reflected}
		reflectedProb = 1.0
	}

	if rand.Float64() < reflectedProb {
		*scattered = Ray{rec.p, reflected}
	} else {
		*scattered = Ray{rec.p, refracted}
	}
	return true
}
