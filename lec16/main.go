package main

import (
	"fmt"
	"log"

	"github.com/syfulsharif/arithmaticGo/helpers"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "MyType"
	myVar.TypeNum = 50

	fmt.Println("myVar Name", myVar.TypeName, "myVar TypeNum", myVar.TypeNum)
}
