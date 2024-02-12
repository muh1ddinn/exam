package main

import (
	"fmt"
)

type People struct {
	Name  string
	Sex   string
	Class string
}

func main() {

	peoplee := []People{}
	var n int
	fmt.Println("enter cout of people ")
	fmt.Scan(&n)

	for i := 0; i <= n; i++ {

		var people = People{}
		fmt.Println("enter %v people's name firs latter big like this Ali : \n", i)
		fmt.Scanln(&people.Name)
		fmt.Println("enter %v people's sex  firs latter big like this MAle: \n", i)
		fmt.Scanln(&people.Sex)
		fmt.Println("enter %v people's job also this : \n", i)
		fmt.Scanln(&people.Class)

		peoplee = append(peoplee, people)
	}

	var choose string
	fmt.Println("You can choose: Male, Female, All")
	fmt.Println("Please enter your choice:")
	fmt.Scanln(&choose)

	var filteredPeople []People
	for _, person := range peoplee {
		switch choose {
		case "Male", "Female":
			if person.Sex == choose {
				filteredPeople = append(filteredPeople, person)
			}
		case "All":
			filteredPeople = append(filteredPeople, person)
		default:
			fmt.Println("Invalid choice. Please choose from Male, Female, or All.")
		}
	}

	fmt.Println(filteredPeople)
}
