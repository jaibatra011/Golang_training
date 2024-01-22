package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "Hello World!, hello go!"
	regex := "ell"

	isPresent := strings.Contains(input, regex) // check if substring is present or not
	if isPresent {
		fmt.Println("The above regex is present in the input")
	} else {
		fmt.Println("The above regex is not present in the input")
	}

	index := strings.Index(input, regex) // index of first match
	fmt.Printf("The above regex is present at index %d \n", index)

	noOfMatches := strings.Count(input, regex) // No of matches present
	fmt.Printf("%d matches are present \n", noOfMatches)

	output := strings.ReplaceAll(input, regex, "*") // replacing all matches with *
	fmt.Println("Old and New strings are " + input + " and " + output)
}
