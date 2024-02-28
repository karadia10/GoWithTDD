package shapes

import "math"

type Shape interface{
  Area() float64
}

type Rectangle struct {
  w, h float64
}

type Circle struct {
  r float64
}

type Triangle struct {
  b,h float64
}

func Perimeter(rect Rectangle) float64 {
  return 2*(rect.w + rect.h)
}

func (rect Rectangle) Area() float64 {
  return rect.w * rect.h
}

func (c Circle) Area() float64 {
  return math.Pi*c.r*c.r
}

func (t Triangle) Area() float64 {
  return 0.5*t.b*t.h
}
