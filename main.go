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

func randomInUnitSphere() Vec3 {
	var p Vec3
	for {
		p = Vec3{rand.Float64(), rand.Float64(), rand.Float64()}.ScalarMultiple(2.0).Substract(Vec3{1, 1, 1})
		if p.SquaredLength() < 1 {
			break
		}
	}
	return p
}

func color(r Ray, world *HitableList) Vec3 {
	var rec HitRecord
	if world.hit(&r, 0.0001, math.MaxFloat64, &rec) {
		target := rec.p.Add(rec.normal).Add(randomInUnitSphere())
		return color(Ray{rec.p, target.Substract(rec.p)}, world).ScalarMultiple(0.5)
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

	nx := 400
	ny := 200
	ns := 100
	var header bytes.Buffer
	fmt.Fprintf(&header, "P3\n%d %d\n255\n", nx, ny)

	_, err = f.WriteString(header.String())
	check(err)
	// Hitable
	list := make([]Hitable, 2)
	list[0] = Sphere{Vec3{0, 0, -1}, 0.5}
	list[1] = Sphere{Vec3{0, -100.5, -1}, 100}
	world := HitableList{list}
	cam := NewCamera()
	fmt.Printf("CAMERA %v\n", cam)
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			col := Vec3{0.0, 0.0, 0.0}
			for s := 0; s < ns; s++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)
				r := cam.GetRay(u, v)
				col = col.Add(color(r, &world))
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
