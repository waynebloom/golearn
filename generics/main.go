package main

import "fmt"

func main() {
  result := add(5, 10)
  fmt.Println(result)
}

func add[T int|float64|string](a, b T) T {
  return a + b
}
