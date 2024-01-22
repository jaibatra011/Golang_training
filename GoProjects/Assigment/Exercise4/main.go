package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//var wg *sync.WaitGroup

	wg := new(sync.WaitGroup)
	wg.Add(2)

	go integers(wg)
	go alphabets(wg)

	wg.Wait()
}

func integers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 6; i++ {
		fmt.Println(i)
		time.Sleep(250 * time.Millisecond)
	}
}

func alphabets(wg *sync.WaitGroup) {
	defer wg.Done()
	alpha := []string{"a", "b", "c", "d", "e"}
	for _, j := range alpha {
		fmt.Println(j)
		time.Sleep(400 * time.Millisecond)
	}
}
