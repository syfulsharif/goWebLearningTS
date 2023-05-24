package main

import "fmt"

func main() {
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	//When I tried Loop throgh the following Slice of animals
	// animals := []string{"dog", "fish", "cow", "goat", "emo"}
	// animals := make(map[string]string)

	// animals["dog"] = "Fido"
	// animals["cat"] = "Mini"
	// animals["horse"] = "Duke"
	// animals["bunny"] = "Fluffy"

	// The following Code will give output of index number and the item
	// for i, animal := range animals {
	// 	fmt.Println(i, animal)
	// }

	//This shows only the slice Item
	// for _, animal := range animals {
	// 	fmt.Println(animal)
	// }
	// Looping through a Map
	// for animalType, animal := range animals {
	// 	fmt.Println(animalType, animal)
	// }

	// var firstLine = "One Upon a midnight dreary"

	// for i, l := range firstLine {
	// 	fmt.Println(i, ":", l)
	// }
	type User struct {
		FirstName string
		LastName  string
		Email     string
		age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john.smith@gmail.com", 40})
	users = append(users, User{"Ethan", "Hunt", "ehunt@gmail.com", 35})
	users = append(users, User{"Elvis", "Menon", "menone@gmail.com", 29})
	users = append(users, User{"Michelle", "Corey", "mcorey@gmail.com", 16})

	for _, l := range users {
		fmt.Println(l.FirstName, l.LastName, l.Email, l.age)
	}

}
