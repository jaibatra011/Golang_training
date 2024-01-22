// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

var wg *sync.WaitGroup

func main() {
	wg = &sync.WaitGroup{}

	mych := make(chan string)
	wg.Add(1)
	//go greet(mych, wg)

	go func(ch chan string) {
		fmt.Println("Greeter")
		ch <- "Hi Amit"
		close(ch)
		fmt.Println("Closed")
		wg.Done()
	}(mych)

	wg.Wait()
}

func greet(ch chan string, wg *sync.WaitGroup) {
	fmt.Println("Greeter")
	ch <- "Hi Amit"
	close(ch)
	fmt.Println("Closed")
	wg.Done()
}
