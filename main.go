package main

import(
  "fmt"
)

func Sum(x int, y int) int {
  return x + y
}

func main() {
  fmt.Println(fmt.Sprintf("Hello World, 5+5 is %d", Sum(5, 5)))
}
