package main

import(
  "fmt"
  "runtime"
)

func Sum(x int, y int) int {
  return x + y
}

func main() {
  fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("GOOS: %s\n", runtime.GOOS)
	fmt.Printf("GOARCH: %s\n", runtime.GOARCH)
  fmt.Printf("Hello World, 5+5 is %d\n", Sum(5, 5))
}
