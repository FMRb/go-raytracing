package main

import "math"

// Vec3 Implementation of a vector
type Vec3 struct {
	X float64
	Y float64
	Z float64
}

// R represent red color on vector
func (v *Vec3) R() float64 {
	return v.X
}

// G represent green color on vector
func (v *Vec3) G() float64 {
	return v.Y
}

// B represent blue color on vector
func (v *Vec3) B() float64 {
	return v.Z
}

// Length of the vector
func (v Vec3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// SquaredLength of the vector
func (v Vec3) SquaredLength() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// UnitVector generation
func (v Vec3) UnitVector() Vec3 {
	return v.ScalarDivide(v.Length())
}

// MakeUnitVector convect vector into a unit vector
func (v *Vec3) MakeUnitVector() {
	k := 1.0 / v.Length()
	v.X *= k
	v.Y *= k
	v.Z *= k
}

// Add two vectors
func (v Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

// Substract two vectors
func (v Vec3) Substract(v2 Vec3) Vec3 {
	return Vec3{v.X - v2.X, v.Y - v2.Y, v.Z - v2.Z}
}

// Multiple two vectors
func (v Vec3) Multiple(v2 Vec3) Vec3 {
	return Vec3{v.X * v2.X, v.Y * v2.Y, v.Z * v2.Z}
}

// Divide two vectors
func (v Vec3) Divide(v2 Vec3) Vec3 {
	return Vec3{v.X / v2.X, v.Y / v2.Y, v.Z / v2.Z}
}

// ScalarMultiple between a value and vector
func (v Vec3) ScalarMultiple(t float64) Vec3 {
	return Vec3{t * v.X, t * v.Y, t * v.Z}
}

// ScalarDivide between a value and vector
func (v Vec3) ScalarDivide(t float64) Vec3 {
	return Vec3{v.X / t, v.Y / t, v.Z / t}
}

// Dot multiplication between two vectors
func (v Vec3) Dot(v2 Vec3) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

// Cross operation on vectors
func (v Vec3) Cross(v2 Vec3) Vec3 {
	return Vec3{
		v.Y*v2.Z - v.Z*v2.Y,
		-(v.X*v2.Z - v.Z*v2.X),
		v.X*v2.Y - v.Y*v2.X}
}
