package dict

import (
	"testing"
)

func TestSearch(t *testing.T) {
  
  dict:=Dictionary{"test": "new test"}
  t.Run("Search in map", func(t *testing.T) {
    got, _ := dict.Search("test")
    want:="new test"
    assertString(t, got, want)  
  })
  t.Run("Unknown word", func(t *testing.T) {
    _, err := dict.Search("Hello")
    assertError(t, err, ErrNotFound)

  })
  t.Run("add word", func(t *testing.T) {
    dict.Add("hello", "world")
    got, err := dict.Search("hello")
    assertNoError(t, err)
    assertString(t, got, "world")
  })
  t.Run("add same word", func(t *testing.T) {
    err:=dict.Add("hello", "world")
    assertError(t, err, ErrWordExist)
  })
  t.Run("update word", func(t *testing.T) {
    err:=dict.Update("hello", "world")
    assertNoError(t, err)
    val, _:=dict.Search("hello")
    assertString(t,  val, "world")
  })
  t.Run("update non existing word", func(t *testing.T) {
    err:=dict.Update("key", "world")
    assertError(t, err, ErrNotFound)
  })
  t.Run("delete existing word", func(t *testing.T) {
    err:=dict.Delete("hello")
    assertNoError(t, err)
  })
  t.Run("delete non existing word", func(t *testing.T) {
    err:=dict.Delete("hello")
    assertError(t, err, ErrNotFound)
  })



}

func assertString(t testing.TB, got, want string) {
  t.Helper()
  if got != want {
    t.Errorf("got %q want %q", got, want)
  }

}

func assertError(t testing.TB, err, want error) {
  t.Helper()
  if err != want {
    t.Errorf("expected %q, got %q", err, want)
  }
  
}

func assertNoError(t testing.TB, err error) {
  t.Helper()
  if err != nil {
    t.Errorf("Error not expected but got : %q", err)
  }
  
}

