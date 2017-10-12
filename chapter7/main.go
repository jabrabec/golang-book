package main

// import (
//  "fmt"
//  "math"
// )

// type Circle struct {
//  x, y, r float64
// }
// type Rectangle struct {
//  x1, y1, x2, y2 float64
// }

// func distance(x1, y1, x2, y2 float64) float64 {
//  a := x2 - x1
//  b := y2 - y1
//  return math.Sqrt(a*a + b*b)
// }

// // func circleArea(x, y, r float64) float64 {
// //     return math.Pi * r * r
// // }
// // func circleArea(c Circle) float64 {
// //     return math.Pi * c.r * c.r
// // }
// // func circleArea(c *Circle) float64 {
// //  return math.Pi * c.r * c.r
// // }
// func (c *Circle) area() float64 {
//  return math.Pi * c.r * c.r
// }
// func (r *Rectangle) area() float64 {
//  l := distance(r.x1, r.y1, r.x1, r.y2)
//  w := distance(r.x1, r.y1, r.x2, r.y1)
//  return l * w
// }
// func main() {
//  // Alternatives:
//  // var c Circle
//  // c:= new(Circle)
//  // c := &Circle{0, 0, 5}
//  // c := Circle{x: 0, y: 0, r: 5}
//  c := Circle{0, 0, 5}
//  r := Rectangle{0, 0, 10, 10}

//  // fmt.Println(c)
//  // fmt.Println(c.x, c.y, c.r)
//  // fmt.Println(circleArea(c))
//  // fmt.Println(circleArea(&c))
//  fmt.Println("area of circle:", c.area())
//  fmt.Println("area of rectangle:", r.area())
// }

// import (
//  "fmt"
// )

// type Person struct {
//  Name string
// }
// type Android struct {
//  // // This way says an Android *has* a person:
//  // Person Person
//  // This way is better; says an Android *is* a person, by way of embedded
//  // types, aka anonymous fields:
//  Person
//  Model string
// }

// func (p *Person) Talk() {
//  fmt.Println("Hi, my name is", p.Name)
// }

// func main() {
//  p := Person{"KITTY"}
//  p.Talk()
//  a := new(Android)
//  a.Person.Talk()
//  b := Android{Person{"Hal"}, "9000"}
//  b.Person.Talk()
// }

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

// // Can handle any number of circles but can't be similarly expanded to include
// // shapes of other types; e.g (circles ...Circle, rectangle ...Rectangles)
// // is invalid:
// func totalArea(circles ...Circle) float64 {
//     var total float64
//     for _, c := range circles {
//         total += c.area()
//     }
//     return total
// }

// // This is non-ideal because arguments are too specific; new shapes have to be
// // added one-by-one:
// func totalArea(circles []Circle, rectangles []Rectangle) float64 {
//     var total float64
//     for _, c := range circles {
//         total += c.area()
//     }
//     for _, r := range rectangles {
//         total += r.area()
//     }
//     return total
// }

type Circle struct {
	x, y, r float64
}
type Rectangle struct {
	x1, y1, x2, y2 float64
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}
func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func (r *Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return 2 * (l + w)
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}
func main() {
	// c := Circle{0, 0, 5}
	// r := Rectangle{0, 0, 10, 10}
	// fmt.Println(c.area())
	// fmt.Println(r.area())
	// fmt.Println(totalArea(&c, &r))
	// multishape := MultiShape{
	//  shapes: []Shape{
	//      Circle{0, 0, 5},
	//      Rectangle{0, 0, 10, 10},
	//  },
	// }
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(c.perimeter())
	fmt.Println(r.perimeter())
}

// Exercises: adding perimeter functions above.
