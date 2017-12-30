package main

import "fmt"

// Fibonacci es una función que devuelve una
// función que devuelve un int.

func fibonacci() func() int {
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
