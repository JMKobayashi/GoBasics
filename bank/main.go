package main

import (
	"fmt"
	"mypackages/numbers"
)

func main() {
	var x uint = 15

	fmt.Printf("The number %d is even? %t \n", x, numbers.Even(x))
}
