package main

import (
	"fmt"
	"os"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	a, b, n := 0, 1, 5

	go func() {
		for i := 0; i < n; i++ {
			if i == 0 {
				ch <- a
			} else if i == 1 {
				ch <- b
			} else {
				ch <- a + b
				a, b = b, a+b
			}
		}
		quit <- true
	}()

	for {
		select {
		case num := <-ch:
			fmt.Println(num)

		case <-quit:
			fmt.Println("All numbers of fibonacci series have been printed")
			os.Exit(0)
		}
	}
}
