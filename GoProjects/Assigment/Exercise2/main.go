package main

import (
	"fmt"
	"strconv"
)

func main() {
	fav := make(map[string][]string)
	fav["bond_James"] = []string{"Shaken,not stirred", "Martinis", "Women"}
	fav["moneypenny_miss"] = []string{"James Bond", "literature", "Computer Science"}
	fav["no_dr"] = []string{"Being evil", "Ice cream", "Sunsets"}

	for key, value := range fav {
		fmt.Println("Favorite things of " + key + " are :")
		for i, j := range value {
			fmt.Println(strconv.Itoa(i+1) + ".) " + j)
		}
		fmt.Println()
	}
}
