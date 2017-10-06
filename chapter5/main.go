package main

import "fmt"

// func main() {
//  var x [5]int
//  x[4] = 100
//  fmt.Println(x)
// }

// func main() {
//     var x [5]float64
//     x[0] = 98
//     x[1] = 93
//     x[2] = 77
//     x[3] = 82
//     x[4] = 83

//     var total float64 = 0
//     for i := 0; i < 5; i++ {
//         total += x[i]
//     }
//     fmt.Println(total / 5)
// }

// func main() {
//     var x [5]float64
//     x[0] = 98
//     x[1] = 93
//     x[2] = 77
//     x[3] = 82
//     x[4] = 83

//     var total float64 = 0
//     for i := 0; i < len(x); i++ {
//         total += x[i]
//     }
//     fmt.Println(total / float64(len(x)))
// }

// func main() {
//  // x := [5]float64{98, 93, 77, 82, 83}

//  // Alternate:
//  // x := [5]float64{
//  //     98,
//  //     93,
//  //     77,
//  //     82,
//  //     83,
//  // }

//  // Another alternate:
//  // If you leave the len declaration = 5 but only have 4 elements, the 5th
//  // element is filled in with a zero value. D:
//  // x := [5]float64{
//  //  98,
//  //  93,
//  //  77,
//  //  82,
//  //  // 83,
//  // }
//  x := [4]float64{
//      98,
//      93,
//      77,
//      82,
//      // 83,
//  }

//  var total float64 = 0
//  for _, value := range x {
//      total += value
//  }
//  fmt.Println(total / float64(len(x)))
// }

// func main() {
//  slice1 := []int{1, 2, 3}
//  slice2 := append(slice1, 4, 5)
//  fmt.Println(slice1, slice2)
// }

// func main() {
//  slice1 := []int{1, 2, 3}
//  slice2 := make([]int, 2)
//  fmt.Println(slice1, slice2)
//  copy(slice2, slice1)
//  fmt.Println(slice1, slice2)
// }

// func main() {
//  x := make(map[string]int)
//  x["key"] = 10
//  fmt.Println(x["key"])
// }

// func main() {
//  elements := make(map[string]string)
//  elements["H"] = "Hydrogen"
//  elements["He"] = "Helium"
//  elements["Li"] = "Lithium"
//  elements["Be"] = "Beryllium"
//  elements["B"] = "Boron"
//  elements["C"] = "Carbon"
//  elements["N"] = "Nitrogen"
//  elements["O"] = "Oxygen"
//  elements["F"] = "Fluorine"
//  elements["Ne"] = "Neon"

//  // name, ok := elements["F"]
//  // fmt.Println(name, ok)
//  // name, ok := elements["Un"]
//  // fmt.Println(name, ok)
//  // if name, ok := elements["F"]; ok {
//  //  fmt.Println(name, ok)
//  if name, ok := elements["Un"]; ok {
//      fmt.Println(name, ok)
//  }
// }

// func main() {
//  elements := map[string]map[string]string{
//      // Example used the following but gofmt removed unnecessary type declarations:
//      // "H": map[string]string{
//      "H": {
//          "name":  "Hydrogen",
//          "state": "gas",
//      },
//      "He": {
//          "name":  "Helium",
//          "state": "gas",
//      },
//      "Li": {
//          "name":  "Lithium",
//          "state": "solid",
//      },
//      "Be": {
//          "name":  "Beryllium",
//          "state": "solid",
//      },
//      "B": {
//          "name":  "Boron",
//          "state": "solid",
//      },
//      "C": {
//          "name":  "Carbon",
//          "state": "solid",
//      },
//      "N": {
//          "name":  "Nitrogen",
//          "state": "gas",
//      },
//      "O": {
//          "name":  "Oxygen",
//          "state": "gas",
//      },
//      "F": {
//          "name":  "Fluorine",
//          "state": "gas",
//      },
//      "Ne": {
//          "name":  "Neon",
//          "state": "gas",
//      },
//  }

//  if el, ok := elements["Li"]; ok {
//      fmt.Println(el["name"], el["state"])
//  }
// }

// // Find smallest number in list:
// func main() {
//  x := []int{
//      48, 96, 86, 68,
//      57, 82, 63, 70,
//      37, 34, 83, 27,
//      19, 97, 9, 17,
//  }
//  y := x[0]

//  for i := 0; i < len(x); i++ {
//      if x[i] <= y {
//          y = x[i]
//      }
//  }
//  fmt.Println(y)
// }

// Official answer:
func main() {
	var min int
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	for i, v := range x {
		if i == 0 || v < min {
			min = v
		}
	}
	fmt.Println(min)
}
