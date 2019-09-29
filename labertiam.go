package main

// Labertiam diffuse algorithm
type Labertiam struct {
	albedo Vec3
}

func (l Labertiam) scatter(rIn *Ray, rec *HitRecord, attenuation *Vec3, scattered *Ray) bool {
	target := rec.p.Add(rec.normal).Add(RandomInUnitSphere())
	*scattered = Ray{rec.p, target.Substract(rec.p)}
	*attenuation = l.albedo
	return true
}
