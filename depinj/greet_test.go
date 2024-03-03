package depinj

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
  t.Run("dep injection printf", func(t *testing.T) {
    buffer:=bytes.Buffer{}
    Greet(&buffer, "Vikas")
    got:=buffer.String()
    want:="Hello, Vikas"
    if got != want {
      t.Errorf("got : %q, want : %q", got, want)
    }
  })
}
