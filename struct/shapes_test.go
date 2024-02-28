package shapes

import "testing"

func TestShape(t *testing.T) {
  got := Perimeter(Rectangle{w:10.0, h:10.0})
  want:=40.0

  if got != want {
    t.Errorf("want %.2f got %.2f", want, got)
  }
}

func TestArea(t *testing.T) {
  checkArea:= func (t testing.TB, shape Shape, want float64) {
    t.Helper()
    got := shape.Area()
    if got != want {
      t.Errorf("want %.2f got %.2f", want, got)
    }

  }
  shapes:=[]struct{
    name string
    shape Shape
    want float64
  }{
    {"rectangle",Rectangle{w:12.0, h: 6.0}, 72.0},
    {"circle",Circle{10}, 314.1592653589793},
    {"triangle",Triangle{12,6}, 36.0},
  }
  for _, tt:=range shapes {
    t.Run(tt.name, func(t *testing.T) {
      checkArea(t, tt.shape, tt.want)
    })
  }
}
