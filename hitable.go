package main

// HitRecord struct for hitable objects
type HitRecord struct {
	t        float64
	normal   Vec3
	p        Vec3
	material Material
}

// Hitable interface
type Hitable interface {
	hit(r *Ray, tMin float64, tMax float64, rec *HitRecord) bool
}
