package array

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
  t.Run("sum of array", func(t *testing.T) {
    arr:=[]int{10,20,30}
    sum:=Sum(arr)
    expected:=60
    if sum != expected {
      t.Errorf("expected %d but got %d, numbers %v", expected, sum, arr)
    }
  })
}

func TestSumAll(t *testing.T) {
  t.Run("summing all array to return array with sum", func(t *testing.T) {
    sumArr:=SumAll([]int{1,2,3}, []int{1,2,3})
    expected:=[]int{6,6}
    if !slices.Equal(sumArr, expected) {
      t.Errorf("expected %v got %v", expected, sumArr)
    }
  })
}

func TestSumAllTails(t *testing.T) {
  t.Run("summing all array tails to return array with sum", func(t *testing.T) {
    sumArr:=SumAllTails([]int{1,2,3}, []int{1,2,3})
    expected:=[]int{5,5}
    if !slices.Equal(sumArr, expected) {
      t.Errorf("expected %v got %v", expected, sumArr)
    }
  })
  t.Run("summing all array with one empty array", func(t *testing.T) {
    sumArr:=SumAllTails([]int{}, []int{1,2,3})
    expected:=[]int{0,5}
    if !slices.Equal(sumArr, expected) {
      t.Errorf("expected %v got %v", expected, sumArr)
    }
  })

}

func BenchmarkSumAll(t *testing.B) {
  for i:=0;i<t.N;i++ {
    sumArr:=SumAll([]int{}, []int{})
    expected:=[]int{}
    if !slices.Equal(sumArr, expected) {
      t.Errorf("expected %v, got %v", expected, sumArr)
    }
  }
}
