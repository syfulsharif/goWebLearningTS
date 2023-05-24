package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Syful",
		LastName:    "Sharif",
		PhoneNumber: "+880 188 88888",
	}

	log.Println(user.FirstName, user.LastName, "Birthdate", user.BirthDate, "Phone", user.PhoneNumber)
}
