package main

import (
	"fmt"
)

// func average(xs []float64) float64 {
//  total := 0.0
//  for _, v := range xs {
//      total += v
//  }
//  return total / float64(len(xs))
// }

// func main() {
//  xs := []float64{98, 93, 77, 82, 83}
//  fmt.Println(average(xs))
// }

// func f() (int, int) {
//  return 5, 6
// }

// func main() {
//  x, y := f()
//  fmt.Println(x, y)
// }

// func add(args ...int) int {
//  total := 0
//  for _, v := range args {
//      total += v
//  }
//  return total
// }

// func main() {
//  // fmt.Println(add(1, 2, 3))
//  xs := []int{1, 2, 3}
//  fmt.Println(add(xs...))
// }

// func makeEvenGenerator() func() uint {
//  i := uint(0)
//  return func() (ret uint) {
//      ret = i
//      i += 2
//      return
//  }
// }
// func main() {
//  nextEven := makeEvenGenerator()
//  fmt.Println(nextEven()) // 0
//  fmt.Println(nextEven()) // 2
//  fmt.Println(nextEven()) // 4
// }

// func factorial(x uint) uint {
//  if x == 0 {
//      return 1
//  }
//  return x * factorial(x-1)
// }
// func main() {
//  fmt.Println(factorial(5))
// }

// func main() {
//  defer func() {
//      str := recover()
//      fmt.Println(str)
//  }()
//  panic("PANIC")
// }

// func zero(x int) {
//  x = 0
// }

// func main() {
//  x := 5
//  fmt.Println(x) // x is 5
//  zero(x)
//  fmt.Println(x) // x is still 5
// }

// func zero(xPtr *int) {
//  *xPtr = 0
// }
// func main() {
//  x := 5
//  fmt.Println(x) // x is 5
//  zero(&x)
//  fmt.Println(x) // x is 0
// }

// func one(xPtr *int) {
//  *xPtr = 1
// }
// func main() {
//  xPtr := new(int)
//  fmt.Println(*xPtr) // x is 0
//  one(xPtr)
//  fmt.Println(*xPtr) // x is 1
// }

// Exercises:
//
// func sum(xs []int) int

// func half(x int) (int, bool) {
//  return x / 2, x%2 == 0
// }
// func main() {
//     fmt.Println(half(2))
// }

// func findMax(args ...int) int {
//  var max int
//  for i, x := range args {
//      if i == 0 || x > max {
//          max = x
//      }
//  }
//  return max
// }
// func main() {
//  fmt.Println(findMax(1, 2, 3, 4, 5, 2))
// }

// func makeOddGenerator() func() uint {
//  i := uint(1)
//  return func() (ret uint) {
//      ret = i
//      i += 2
//      return
//  }
// }
// func main() {
//  nextOdd := makeOddGenerator()
//  fmt.Println(nextOdd()) // 1
//  fmt.Println(nextOdd()) // 3
//  fmt.Println(nextOdd()) // 5
// }

// func fib(n int) int {
// 	switch n {
// 	case 0:
// 		return 0
// 	case 1:
// 		return 1
// 	default:
// 		return fib(n-1) + fib(n-2)
// 	}
// }
// func main() {
// 	fmt.Println(fib(1))
// }

// // Recover from a runtime panic:
// func main() {
// 	defer func() {
// 		str := recover()
// 		fmt.Println(str)
// 	}() // anonymous function call required or else panic is never executed
// 	panic("PANIC!")
// }

// // Get mem address of variable:
// x = (&y)
// // Assign value to pointer:
// *x = 5
// // Create a new pointer:
// x = new(int)

// // Value of x for this program:
// func square(x *float64) {
// 	*x = *x * *x
// }
// func main() {
// 	x := 1.5
// 	square(&x) // x is 2.25
// }

// Program to swap two integers:
// (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1)
func swap(xPtr *int, yPtr *int) {
	*xPtr, *yPtr = *yPtr, *xPtr
}
func main() {
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println(x, y)
}

// // Official answer:
// func swap(x, y *int) {
//     *x, *y = *y, *x
// }
