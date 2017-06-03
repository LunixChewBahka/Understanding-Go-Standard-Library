/*

Package sort provides primitives for sorting slices and user-defined collections.

*/
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on the Age field.
type ByAge []Person
type ByLength []string

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (s ByLength) Len() int           { return len(s) }
func (s ByLength) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByLength) Less(i, j int) bool { return len(s[i]) < len(s[j]) }

func main() {
	// var person takes the supplied values stored inside the []Person container
	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
		{"Jar", 26},
	}

	fruits := []string{"peach", "banana", "kiwi"}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)

	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}
