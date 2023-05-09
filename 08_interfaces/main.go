package main

import (
	"fmt"
)

// Interfaces
// An interface is a set of method signatures.
// A type implements an interface by implementing its methods.
// There is no explicit declaration of intent, no "implements" keyword.

type Human interface {
	breath()
	walk()
}

type AnotherInterface interface {
	breath()
	walk()
}

// Adult and child types implement the Human interface.
type Adult struct {
	name string
}

func (a Adult) breath() {
	println(a.name, "is breathing.")
}

func (a Adult) walk() {
	println(a.name, "is walking.")
}

// Function that takes an interface as an argument
func beingHuman(h Human) {
	h.breath()
	fmt.Println("I am being human.")
}

type Child struct {
	name string
	age int
}

func (c Child) breath() {
	println("Child", c.name, "is breathing.")
}

func main() {
	a := Adult{name: "Jane"}
	a.walk()
	beingHuman(a)

	c := Child{name: "Tom", age: 5}
	c.breath()
}
