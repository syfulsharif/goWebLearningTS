package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]User)

	me := User{
		FirstName: "Syful",
		LastName:  "Sharif",
	}

	myMap["me"] = me

	log.Println(myMap["me"].FirstName)

}
