package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Bread string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Milou",
		Bread: "Terier",
	}

	gorilla := Gorilla {
		Name: "Mangani",
		Color: "Gray",
		NumberOfTeeth: 27,
	}

	PrintInfo(&dog)

	PrintInfo(&gorilla)

}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(),"legs")
}

func (d *Dog) Says() string {
	return "Woof"
	
}

func (d *Dog) NumberOfLegs() int  {
	return 4
	
}

func (d *Gorilla) Says() string {
	return "Ugh"
	
}

func (d *Gorilla) NumberOfLegs() int  {
	return 2
	
}