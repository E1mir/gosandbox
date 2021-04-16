package main

import (
	"fmt"
	"math/rand"
	"time"
)

var x = 10
var y = "Hello World!"
var z int

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	a := 'F'
	b := rand.Float64() > 0.5
	c := 100
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
