package main

// Metal material representing a mirror metal
type Metal struct {
	albedo Vec3
	fuzz   float64
}

// NewMetal create a new Metal object
func NewMetal(albedo Vec3, f float64) Metal {
	fuzz := 1.0
	if f < 1 {
		fuzz = f
	}
	return Metal{albedo, fuzz}
}

func (m Metal) scatter(rIn *Ray, rec *HitRecord, attenuation *Vec3, scattered *Ray) bool {
	reflected := Reflect(rIn.Direction().UnitVector(), rec.normal)
	*scattered = Ray{rec.p, reflected.Add(RandomInUnitSphere().ScalarMultiple(m.fuzz))}
	*attenuation = m.albedo
	return scattered.Direction().Dot(rec.normal) > 0
}
