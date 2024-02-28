package main
import "testing"

func TestHello(t* testing.T) {
  t.Run("saying Hello to people", func(t * testing.T) {
    got:= Hello("Vikas", "")
    want:="Hello, Vikas"
     assertCorrectMessage(t, got, want)
  })
  t.Run("saying Hello, World when empty string is provided", func(t *testing.T) {
    got:=Hello("", "")
    want:="Hello, world"
    assertCorrectMessage(t, got, want)
  })
  t.Run("in spanish", func(t *testing.T) {
    got := Hello("vikas", "spanish")
    want := "Hola, vikas"
    assertCorrectMessage(t, got, want)
  })
}

func assertCorrectMessage(t testing.TB, got, want string) {
  t.Helper()
  if got != want{
    t.Errorf("got %q want %q", got, want)
  }
}
