package main

import "log"

func main() {
	for i := 0; i <= 5; i++ {
		log.Println(i)

	}

	animals := []string{"dog", "fish","horse", "cow", "cat"}

	for _ ,animal := range animals {
		log.Println(animal)
		
	}

	animalmap := make(map[string]string)

	animalmap["dog"] = "Milo"
	animalmap["cat"] = "Fluffy"
	for animalType, animal := range animalmap {

		log.Println(animalType, animal)
		
	}

	type User struct {
		FirstName string
		LastName  string
		Email 	  string
		Age		  int	
	}

	var users []User
	users = append(users, User{"Jhon", "Smith", "jhone@gmail.com", 30})
	users = append(users, User{"Tal", "Baror", "tal@gmail.com", 50})
	users = append(users, User{"Sandra", "Hamou", "sandra@gmail.com", 52})
	users = append(users, User{"Noa", "Baror", "noa@gmail.com", 30})

	for _, l := range users{
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
		
	}
}