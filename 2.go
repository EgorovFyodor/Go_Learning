package main

import (
	"fmt"
	"math"
	"math/rand"
)

func Print_1() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Printf("Now you have %g problems.", math.Sqrt(7))
}
