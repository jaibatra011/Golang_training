package main

import (
	"fmt"
)

func main() {
	n := 8
	ans, _ := fibonacci(n)
	fmt.Printf("Fibonacci of %d is : %d", n, ans)
}

func fibonacci(n int) (int, int) {
	if n == 2 {
		return 2, 1
	}
	ans, ansN := fibonacci(n - 1)
	return ans + ansN, ans
}
