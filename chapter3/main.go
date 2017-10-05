package main

import "fmt"

// Convert feet to meters:
func main() {
	fmt.Println("Enter a measurement in feet: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 0.3048

	fmt.Println(output)
}

// // Convert deg F to deg C:
// func main() {
//  fmt.Println("Enter a Fahrenheit temperature: ")
//  var input float64
//  fmt.Scanf("%f", &input)

//  output := (input - 32) * 5 / 9

//  fmt.Println(output)
// }

// func main() {
//  fmt.Println("Enter a number: ")
//  var input float64
//  fmt.Scanf("%f", &input)

//  output := input * 2

//  fmt.Println(output)
// }

// Alternate:
// func main() {
//     var x string
//     x = "Hello, World"
//     fmt.Println(x)
// }
// func main() {
//     var x string = "Hello, World"
//     fmt.Println(x)
// }

// func main() {
//     var x string
//     x = "first"
//     fmt.Println(x)
//     x = "second"
//     fmt.Println(x)
// }
// func main() {
//     var x string
//     x = "first"
//     fmt.Println(x)
//     x = x + "second"
//     fmt.Println(x)
// }
// func main() {
//     var x string
//     x = "first"
//     fmt.Println(x)
//     x += "second"
//     fmt.Println(x)
// }

// func main() {
//     var x string = "hello"
//     var y string = "world"
//     fmt.Println(x == y)
// }
// func main() {
//     var x string = "hello"
//     var y string = "hello"
//     fmt.Println(x == y)
// }

// func main() {
//     // x := "Hello, World"
//     // var x = "Hello, World"
//     x := 5
//     fmt.Println(x)
// }

// func main() {
//     dogsName := "Max"
//     fmt.Println("My dog's name is", dogsName)
// }

// func main() {
//  const x string = "Hello, World"
//  fmt.Println(x)
// }
