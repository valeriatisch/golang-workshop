package main

import "fmt"

// Structs
// - A struct is a collection of fields.
// - A struct is declared by specifying the name of the struct and the type of each field.
type Person struct {
	Name string
	age  int
}

// Embedded structs
// - A struct can contain other structs.
type Student struct {
	Person
	grade float64
}

func main() {
	// Arrays
	// - An array is a fixed-size collection of elements of the same type.
	// - An array is declared by specifying the type and size of the array.
	// - An array is indexed starting at 0.
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)

	// Length of an array is the number of elements it was declared with.
	fmt.Println(len(arr))

	// Slices
	// - A slice is a dynamically-sized, flexible view into the elements of an array.
	var slice []int
	fmt.Println(len(slice))
	slice = append(slice, 2)
	fmt.Println(slice)
	fmt.Println(len(slice))

	// composite literal
	// x := type{values}
	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice2)

	// multi-dimensional slices
	// - A slice can contain any type, including other slices.
	slice3 := [][]int{{1, 2, 3}, {10, 11, 19}}
	fmt.Println(slice3)

	// Maps
	// - A map is an unordered collection of key-value pairs.
	// - A map is declared by specifying the type of its key and value.
	// - A map is indexed by its key.
	m := map[string]int{
		"foo": 42,
		"bar": 69}
	fmt.Println(m)
	fmt.Println(m["foo"])
	fmt.Println(m["kevin"])

	// No entry
	val, ok := m["kevin"]
	fmt.Println(val, ok)
	val2, ok2 := m["foo"]
	fmt.Println(val2, ok2)

	// Adding new element
	m["new"] = 100

	// Deleting an element
	delete(m, "foo")
	fmt.Println(m)

	// Structs
	var p Person = Person{
		Name: "Kevin",
		age:  21,
	}
	fmt.Println(p)

	// Embedded structs
	var s Student = Student{
		Person: Person{
			Name: "Jane",
			age:  20,
		},
		grade: 2.0,
	}
	fmt.Println(s)

	// Anonymous struct
	anonym := struct {
		name  string
		color string
	}{
		name:  "John",
		color: "blue",
	}
	fmt.Println(anonym)

}
