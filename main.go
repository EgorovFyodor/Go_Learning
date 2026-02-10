package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World")
	Prnt()
	for i := 0; i < 3; i++ {
		Tm()
	}

	fmt.Println(Sqrt(2))
}
