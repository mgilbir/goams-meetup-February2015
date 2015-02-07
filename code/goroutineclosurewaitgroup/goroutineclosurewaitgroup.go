package main

import "fmt"
import "sync"

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Hello world")
		}()
	}
	wg.Wait()
}
