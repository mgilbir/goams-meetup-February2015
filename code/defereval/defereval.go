package main

import "fmt"

func donothing() (err error) {
	// err = fmt.Errorf("error: %s\n", "set before defer")
	defer fmt.Println(err)

	err = fmt.Errorf("error: %s\n", "set after defer")

	return err
}

func main() {
	err := donothing()
	if err != nil {
		fmt.Printf("Back in main... %s", err)
	}
}
