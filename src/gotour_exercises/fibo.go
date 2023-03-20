package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib1, fib2 := 0, 1
	return func() int {
		fib_prev := fib1
		fib1 = fib2
		fib2 += fib_prev
		return fib_prev
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
