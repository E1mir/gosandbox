package main

import "fmt"

func main() {
	a, b := 2, 5
	fmt.Println("Hello world")
	fmt.Printf("%d + %d = %d", a, b, add(a, b))
	onlyEven(30)
}

func add(a, b int) int {
	return a + b
}

func onlyEven(intRange int) {
	for i := 0; i < intRange; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
