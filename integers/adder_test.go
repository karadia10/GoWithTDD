package main

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
  t.Run("adding two numbers", func(t *testing.T) {
    sum := Add(2, 2)
    expected:= 4
    if sum != expected {
      t.Errorf("expected %d but got %d", expected, sum)
    }

  })
 }

func ExampleAdd() {
  sum:=Add(1, 5)
  fmt.Println(sum)
  // Output: 6
}
