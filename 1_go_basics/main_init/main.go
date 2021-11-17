package main

import "fmt"

func init() {
	fmt.Printf("This is init #1")
}

func init() {
	fmt.Printf("This is init #2")
}

func main() {
	fmt.Printf("This is main function")
}
