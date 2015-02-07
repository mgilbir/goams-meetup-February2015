package main

import "fmt"

func donothing() {
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")

	fmt.Println("Not deferred")
}

func main() {
	donothing()
}
