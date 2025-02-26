package main

import "fmt"

func MapsInGo() {

	var heroes map[int]string
	heroes = make(map[int]string)

	heroes[1] = "Batman"
	heroes[2] = "Superman"
	fmt.Println(heroes)

	villains := make(map[int]string)
	scores := map[string]int{"Kills": 43, "Damage": 12000}

	fmt.Println(villains)
	fmt.Println(scores)

	fmt.Println(scores["Kills"])
	fmt.Println(heroes[1])

	fmt.Println(heroes[5])     // key has no value then it returns string
	fmt.Println(scores["yOo"]) //in case of int it returns 0 when no val

	boll_check := map[int]bool{1: true, 2: false}
	fmt.Println(boll_check[1])
	fmt.Println(boll_check[5]) // key don't exists hence false is returned for bool

	// other ways to check keys ~ return true or false
	_, hasKey1 := heroes[1]
	fmt.Println(hasKey1)

	_, hasKey2 := heroes[7]
	fmt.Println(hasKey2)

	heroes[3] = "Flash"
	fmt.Println(heroes)
	delete(heroes, 3)

	fmt.Println(heroes)
}
