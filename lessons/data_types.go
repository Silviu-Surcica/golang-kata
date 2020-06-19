package lessons

import "fmt"
import "strings"

// Pointers in Go work exactly as in C but Go has no pointer arithmetic
func Pointers(){
    a := 2
    p := &a //point to a
    fmt.Println(*p)
    *p = *p + 3 //assign value of a through pointer p
    fmt.Println(a)
}

// Go doesnt have classes but has a data type called Struct
func Structs(){
    // definition of a struct
    type Square struct{
        long float64
        lat float64
    }
    s1 := Square{} // long 0, lat 0 . {0, 0}
    fmt.Println(s1)
    s2 := Square{1.0, 2.5}
    fmt.Println(s2)
    s3 := Square{long: 1.0} // long 1.0, lat 0
    fmt.Println(s3)
    // you can also have a pointer to a struct
    p := &Square{}
    // you can access structs'field using dot. The language permits p.lat instead of (*p).lat
    fmt.Println(p.lat)
}

func ArraysAndSlices(){
    // declares an array of length 3 of type bool. Arrays have a fixed size.
    a := [3]bool{true, false, true} //this is an array
    /* Slices are more flexible,they have a dynamic size.
       You can get slice out of an array by specifying two limits:
       s := a[0:2]
       This slice is like a reference to the array. If you modify the slice,
        it will change the array and all other slices that share that array.
    */
    s := []bool{true, false, true} //this is a slice
    fmt.Println(a,s)
    /*
    A slice has both a length and a capacity
    The length of a slice is the number of elements it contains. len(s)
    The capacity of a slice is the number of elements in the underlying array,
    counting from the first element in the slice. cap(s)
    You can create a dynamic sized array using the builtin make function:
    a := make([]int, 5) // len(a) 5, cap(a) 5
    a := make([]int, 0, 5) // len(a) 0, cap(a) 5
    */
    //If the backing array of s is too small to fit all the given values a bigger array will be allocated.
    //The returned slice will point to the newly allocated array.
    s = append(s, true) // builtin function to add new elements to the slice
    s = append(s, true, false, false) // you can append more items

    // the builtin range function can iterate over a slice
    for index, value := range s { // you can ommit the index or the value:
                                  // for _, value := range s {}
                                  // for index, _ := range s {}
                                  // for index := range s {}
        fmt.Println(index, value)
    }
}

// A map maps keys to values, like dict in python
func Maps(){
    var m = map[string]int{"first": 1}
    fmt.Println(m["first"])
    var n = map[int]int{1: 1}
    fmt.Println(n[1])

    type Point struct {
	    X, Y int
    }
    var mapper = map[string]Point{
        "point1": {1, 2},
        "point2": {3, 4}, //this last comma is required
    }
    fmt.Println(mapper)
    /* test that a key is present. exists will be true or false.
    If the key is not in mapper then a will have the zero value for map's element type, which is {0 0} in our case.
    */
    a, exists := mapper["point3"]
    fmt.Println(a, exists)
    // delete an element
    delete(mapper, "point2")

    // You can also use the range function on maps
    for key, value := range mapper {
        fmt.Println(key, value)
    }
    var o = map[string][]string{"first": {"a", "b"}}
    for key, value := range o {
        fmt.Println(key, value)
    }
}

/* ===========EXERCISE===========
Implement the function WordLength
It takes a string and returns a map
with a string key and int as its value.
This function will iterate over each word and will count its letters.
Helper function: https://golang.org/pkg/strings/#Fields
Ex:
WorldLength("A world full of dreamers")
{"A": 1, "world": 5, "full": 4, "of": 2, "dreamers": 8}
*/

func WordLength(s string) map[string]int {
    result := make(map[string]int)
    words := strings.Fields(s)
    for _, word := range words {
        result[word] = len(word)
    }
    return result
}
