package repeat

import "testing"

func TestRepeat(t *testing.T) {
  t.Run("testing string in loop", func(t *testing.T) {
    repeat:=Repeat("a", 5)
    expected:="aaaaa"

  if repeat != expected {
      t.Errorf("Expected %q but got %q", expected, repeat)
    }
  })
}

func BenchmarkRepeat(t *testing.B) {
  for i:=0;i<t.N;i++ {
    Repeat("a", 5)
  }
}
