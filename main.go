package main

import(
  "fmt"
)

func main() {
  fmt.Println(addOneToNumberOfTerms(100)) // 1 + 2 + 3 + 4 ... + 100
  fmt.Println(addSequence(100, 1, 1)) // same as above
  fmt.Println(addSequence(3, 1, 1)) // 1 + 2 + 3
  fmt.Println(addSequence(3, 5, 6)) // 6 + 11 + 16
}

func addOneToNumberOfTerms(numberOfTerms int64) int64 {
  n := numberOfTerms
  return n*(n+1)/2
}

func addSequence(numberOfTerms int64, differenceBetweenTerms int64, firstTerm int64) int64 {
  n := numberOfTerms
  d := differenceBetweenTerms
  a := firstTerm
  return n*(2*a+(n-1)*d)/2
}
