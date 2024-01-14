package basics

import "fmt"

func Arrays() {
	var cities [5]string
	cities[0] = "Moscow"
	cities[1] = "Budapest"
	cities[2] = "Taiwan"
	cities[3] = "Beijing"
	cities[4] = "Halifax"
	fmt.Println(cities)

	for _, city := range cities {
		fmt.Printf("City -> %v\n", city)
	}
	fmt.Println(len(cities))

	users := []string{
		"Mario",
		"Bowser",
	}
	users = append(users, "Tifa")
	users = append(users, "Luigi")
	users = append(users, "Yoshi")

	for _, v := range users {
		fmt.Printf("Users -> %v\n", v)
	}

	verticles := make(map[string]int)

	verticles["Triangle"] = 2
	verticles["Square"] = 3
	verticles["Dodecagon"] = 12

	delete(verticles, "Dodecagon")

	for v, k := range verticles {
		fmt.Printf("Keys -> %v : Values -> %v\n", v, k)
	}
}
