package array

func Sum(arr []int) (sum int) {
  for _, v:= range arr {
    sum+=v
  }
  return
}

func SumAll(arr ...[]int) []int {
  sumArr:=make([]int, 0)
  for _, v:=range arr {
    sumArr = append(sumArr, Sum(v))
  }
  return sumArr
}

func SumAllTails(arr ...[]int) []int {
  sumArr:=make([]int,0)
  for _,v:=range arr {
    if len(v) <= 1 {
      sumArr = append(sumArr, 0)
      continue
    }
    sumArr = append(sumArr, Sum(v[1:]))
  }
  return sumArr
}
