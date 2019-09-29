package main

// Metal material representing a mirror metal
type Metal struct {
	albedo Vec3
}

func reflect(v Vec3, n Vec3) Vec3 {
	return v.Substract(n.ScalarMultiple(2.0 * v.Dot(n)))
}

func (m Metal) scatter(rIn *Ray, rec *HitRecord, attenuation *Vec3, scattered *Ray) bool {
	reflected := reflect(rIn.Direction().UnitVector(), rec.normal)
	*scattered = Ray{rec.p, reflected}
	*attenuation = m.albedo
	return scattered.Direction().Dot(rec.normal) > 0
}
