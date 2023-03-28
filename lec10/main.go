package name

import "time"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName: "Syful",
		LastName:  "Sharif",
	}
}
