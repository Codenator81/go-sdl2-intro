package main

import "math"

type Vector2d struct {
    x float64
    y float64
}

func(v *Vector2d) X() float64 { return v.x }
func(v *Vector2d) Y() float64{ return v.y }
func(v *Vector2d) SetX(x float64) { v.x = x }
func(v *Vector2d) SetY(y float64) { v.y = y }

func(v *Vector2d) Length() float64{ return math.Sqrt(v.x * v.x + v.y * v.y) }

// Plus func
func(v *Vector2d) Add(v1 Vector2d) {
    v.x += v1.x
    v.y += v1.y
}

func(v *Vector2d) Subtr(v1 Vector2d) {
    v.x -= v1.x
    v.y -= v1.y
}

func(v *Vector2d) Multiply(scalar float64) {
    v.x *= scalar
    v.y *= scalar
}

func(v *Vector2d) Divide(scalar float64) {
    v.x /= scalar
    v.y /= scalar
}

func(v *Vector2d) Normalize() {
    l := float64(v.Length())
    if  l > 0 {
        d := float64(1 / l)
        v.Multiply(d)
    }
}

