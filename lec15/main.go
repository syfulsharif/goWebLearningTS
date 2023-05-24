package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Cat struct {
	Name            string
	Breed           string
	AverageLifeTime int
}

func main() {
	dog := Dog{
		Name:  "Tommy",
		Breed: "German Shepered",
	}

	cat := Cat{
		Name:            "Jasper",
		Breed:           "Turkish",
		AverageLifeTime: 12,
	}
	PrintInfo(&dog)
	PrintInfo(&cat)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has ", a.NumberOfLegs(), "number of legs")
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Cat) Says() string {
	return "Meaw"
}

func (d *Cat) NumberOfLegs() int {
	return 4
}
