package main

import "fmt"

func soma(x int) int {
	lista := 0
	for i := 1; i <= x; i++ {
		lista += i
	}
	return lista
}

func main() {
	fmt.Println(soma(22))
}
