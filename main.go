package main

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// RandomInUnitSphere generate a random unit vector
func RandomInUnitSphere() Vec3 {
	var p Vec3
	for {
		p = Vec3{rand.Float64(), rand.Float64(), rand.Float64()}.ScalarMultiple(2.0).Substract(Vec3{1, 1, 1})
		if p.SquaredLength() < 1 {
			break
		}
	}
	return p
}

func color(r Ray, world *HitableList, depth int) Vec3 {
	var rec HitRecord
	if world.hit(&r, 0.0001, math.MaxFloat64, &rec) {
		var scattered Ray
		var attenuation Vec3
		if depth < 50 && rec.material.scatter(&r, &rec, &attenuation, &scattered) {
			return attenuation.Multiple(color(scattered, world, depth+1))
		}
		return Vec3{0, 0, 0}
	}
	unitDirection := r.Direction().UnitVector()
	t := 0.5 * (unitDirection.Y + 1.0)
	// (1.0-t)*vec3{1.0,1.0,1.0} + t*vec3{0.5,0.7,1.0}
	return Vec3{1.0, 1.0, 1.0}.ScalarMultiple(1.0 - t).Add(Vec3{0.5, 0.7, 1.0}.ScalarMultiple(t))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	f, err := os.Create("image.ppm")
	check(err)
	defer f.Close()

	// Size of the viewport
	nx := 400
	ny := 200
	// Number of ray samples per pixel
	ns := 100
	var header bytes.Buffer
	fmt.Fprintf(&header, "P3\n%d %d\n255\n", nx, ny)

	_, err = f.WriteString(header.String())
	check(err)
	// Hitable
	list := make([]Hitable, 5)
	list[0] = Sphere{Vec3{0, 0, -1}, 0.5, Labertiam{Vec3{0.1, 0.2, 0.5}}}
	list[1] = Sphere{Vec3{0, -100.5, -1}, 100, Labertiam{Vec3{0.8, 0.8, 0.0}}}
	list[2] = Sphere{Vec3{1, 0, -1}, 0.5, NewMetal(Vec3{0.8, 0.6, 0.2}, 0.0)}
	list[3] = Sphere{Vec3{-1, 0, -1}, 0.5, Dielectric{1.5}}
	list[4] = Sphere{Vec3{-1, 0, -1}, -0.45, Dielectric{1.5}}
	world := HitableList{list}
	lookFrom := Vec3{3, 3, 2}
	lookAt := Vec3{0, 0, -1}
	distToFocus := lookFrom.Substract(lookAt).Length()
	aperture := 1.0
	cam := NewCamera(lookFrom, lookAt, Vec3{0, 1, 0}, 20,
		float64(nx/ny), aperture, distToFocus)
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			col := Vec3{0.0, 0.0, 0.0}
			for s := 0; s < ns; s++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)
				r := cam.GetRay(u, v)
				col = col.Add(color(r, &world, 0))
			}
			col = col.ScalarDivide(float64(ns))
			col = Vec3{math.Sqrt(col.X), math.Sqrt(col.Y), math.Sqrt(col.Z)}
			ir := int(255.99 * col.R())
			ig := int(255.99 * col.G())
			ib := int(255.99 * col.B())
			var colorLine bytes.Buffer
			fmt.Fprintf(&colorLine, "%d %d %d\n", ir, ig, ib)
			_, err = f.WriteString(colorLine.String())
			check(err)
		}
	}
}
