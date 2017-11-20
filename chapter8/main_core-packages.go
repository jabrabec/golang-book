package main

import (
	// "container/list"
	// "errors"
	"fmt"
	// "io/ioutil"
	"hash/crc32"
	// "os"
	// "path/filepath"
	// "sort"
	// "strings"
)

// String functions:

// func main() {
// 	// Contains(s, substr string) bool
// 	fmt.Println(strings.Contains("test", "es"))
// }

// func main() {
// 	// func Count(s, sep string) int
// 	fmt.Println(strings.Count("test", "t"))
// }

// func main() {
// 	// func HasPrefix(s, prefix string) bool
// 	fmt.Println(strings.HasPrefix("test", "te"))
// }

// func main() {
// 	// func HasSuffix(s, suffix string) bool
// 	fmt.Println(strings.HasSuffix("test", "st"))
// }

// func main() {
// 	// func HasSuffix(s, suffix string) bool
// 	fmt.Println(strings.HasSuffix("test", "st"))
// }

// func main() {
// 	// func Index(s, sep string) int
// 	fmt.Println(strings.Index("test", "e"))
// }

// func main() {
// 	// func Join(a []string, sep string) string
// 	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
// }

// func main() {
// 	// func Repeat(s string, count int) string
// 	fmt.Println(strings.Repeat("a", 5))
// }

// func main() {
//     // func Replace(s, old, new string, n int) string
//     fmt.Println(strings.Replace("aaaa", "a", "b", 2))
// }

// func main() {
// 	// func Split(s, sep string) []string
// 	fmt.Println(strings.Split("a-b-c-d-e", "-"))
// }

// func main() {
//     // func ToLower(s string) string
//     fmt.Println(strings.ToLower("TEST"))
// }

// func main() {
//     // func ToUpper(s string) string
//     fmt.Println(strings.ToUpper("test"))
// }

// func main() {
// 	arr := []byte("test")
// 	str := string([]byte{'t', 'e', 's', 't'})
// 	fmt.Println(arr)
// 	fmt.Println(str)
// }

// Files & Folders

// func main() {
// 	file, err := os.Open("test.txt")
// 	if err != nil {
// 		// handle the error here
// 		return
// 	}
// 	defer file.Close()

// 	// get the file size
// 	stat, err := file.Stat()
// 	if err != nil {
// 		return
// 	}
// 	// read the file
// 	bs := make([]byte, stat.Size())
// 	_, err = file.Read(bs)
// 	if err != nil {
// 		return
// 	}

// 	str := string(bs)
// 	fmt.Println(str)
// }

// func main() {
// 	bs, err := ioutil.ReadFile("test.txt")
// 	if err != nil {
// 		return
// 	}
// 	str := string(bs)
// 	fmt.Println(str)
// }

// func main() {
// 	file, err := os.Create("temp.txt")
// 	if err != nil {
// 		// handle the error here
// 		return
// 	}
// 	defer file.Close()

// 	file.WriteString("test")
// }

// func main() {
// 	dir, err := os.Open(".")
// 	if err != nil {
// 		return
// 	}
// 	defer dir.Close()

// 	// -1 returns all entries in directory, otherwise is a numeric limiter
// 	// of # of entries returned
// 	fileInfos, err := dir.Readdir(-1)
// 	if err != nil {
// 		return
// 	}
// 	for _, fi := range fileInfos {
// 		fmt.Println(fi.Name())
// 	}
// }

// // Folder contents recursively:
// // use filepath.SkipDir to stop walking immediately
// func main() {
// 	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
// 		fmt.Println(path)
// 		return nil
// 	})
// }

// Errors

// create a new error type:
// func main() {
// 	err := errors.New("error message")
// }

// Containers and Sort:

// func main() {
//     // create empty linked list; zero value is []
// 	var x list.List
//     // append to list:
// 	x.PushBack(1)
// 	x.PushBack(2)
// 	x.PushBack(3)

// 	for e := x.Front(); e != nil; e = e.Next() {
// 		fmt.Println(e.Value.(int))
// 	}
// }

// // Implement our own data and types to manipulate; "sort.Interface requires
// // three methods: Len, Less, and Swap"
// type Person struct {
// 	Name string
// 	Age  int
// }

// // type ByName []Person

// // func (ps ByName) Len() int {
// // 	return len(ps)
// // }
// // func (ps ByName) Less(i, j int) bool {
// // 	return ps[i].Name < ps[j].Name
// // }
// // func (ps ByName) Swap(i, j int) {
// // 	ps[i], ps[j] = ps[j], ps[i]
// // }

// // Define our own sort (first example sorted by name):
// type ByAge []Person

// func (this ByAge) Len() int {
// 	return len(this)
// }
// func (this ByAge) Less(i, j int) bool {
// 	return this[i].Age < this[j].Age
// }
// func (this ByAge) Swap(i, j int) {
// 	this[i], this[j] = this[j], this[i]
// }

// func main() {
// 	kids := []Person{
// 		{"Jill", 9},
// 		{"Jack", 10},
// 	}
// 	// sort.Sort(ByName(kids))
// 	sort.Sort(ByAge(kids))
// 	fmt.Println(kids)
// }

// Hashes and Cryptography:

func main() {
	// create a hasher
	h := crc32.NewIEEE()
	// write our data to it
	h.Write([]byte("test"))
	// calc the crc32 checksum
	v := h.Sum32()
	fmt.Println(v)
}
