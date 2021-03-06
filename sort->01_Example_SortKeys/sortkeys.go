package main

import (
	"fmt"
	"sort"
)

// A acouple of type definitions to make the units clear.
type earthMass float64
type au float64

// A Planet defines the properties of a solar system object
type Planet struct {
	name     string
	mass     earthMass
	distance au
}

// By is the type of a "less" function that defines the ordering of it's Planet
// arguments.
type By func(p1, p2 *Planet) bool

// Sort is a method on the function type, By, that sorts the argument slice
// according to the function.
func (by By) Sort(planets []Planet) {
	ps := &planetSorter{
		planets: planets,
		// The Sort method's receiver is the function (closure) that defines
		// the sort order
		by: by,
	}
	sort.Sort(ps)
}

// planetSorter joins a By function and a slice of Planets to be sorted.
type planetSorter struct {
	planets []Planet
	by      func(p1, p2 *Planet) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *planetSorter) Len() int {
	return len(s.planets)
}

// Swap is part of sort.Interface
func (s *planetSorter) Swap(i, j int) {
	s.planets[i], s.planets[j] = s.planets[j], s.planets[i]
}

// Less is part of sort Interface, It is implemented by calling the "by" closure
// in the sorter
func (s *planetSorter) Less(i, j int) bool {
	// we now pass the address of indexe's i and j
	return s.by(&s.planets[i], &s.planets[j])
}

var planets = []Planet{
	// Name, mass and distance
	{"Mercury", 0.0555, 0.4},
	{"Venus", 0.815, 0.7},
	{"Earth", 1.0, 1.0},
	{"Mars", 0.107, 1.5},
	{"Jupiter", 318, 5.2},
	{"Saturn", 95, 9.5},
	{"Uranus", 14, 19.2},
	{"Neptune", 17, 30.1},
}

// ExampleSortKeys demonstrates a technique for sorting a struct type using
// programmable sort criteria
func main() {
	// Closures that order the Planet structure
	name := func(p1, p2 *Planet) bool {
		return p1.name < p2.name
	}

	mass := func(p1, p2 *Planet) bool {
		return p1.mass < p2.mass
	}

	distance := func(p1, p2 *Planet) bool {
		return p1.distance < p2.distance
	}

	decreasingDistance := func(p1, p2 *Planet) bool {
		return !distance(p1, p2)
	}

	// Sort the planets by the various criteria
	By(name).Sort(planets)
	fmt.Println("By name:", planets)

	By(mass).Sort(planets)
	fmt.Println("By mass:", planets)

	By(distance).Sort(planets)
	fmt.Println("By distance:", planets)

	By(decreasingDistance).Sort(planets)
	fmt.Println("By decreasing distance:", planets)
}
