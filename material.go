package main

// Material abstract
// Produce a scatter ray, if scattered how much the ray should be attenuated
type Material interface {
	scatter(rIn *Ray, rec *HitRecord, attenuation Vec3, scattered *Ray) bool
}
