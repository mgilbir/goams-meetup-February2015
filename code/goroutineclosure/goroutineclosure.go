package main

import "fmt"
import "sync"

func main() {
	var wg sync.WaitGroup

	for _, s := range []string{"a", "b", "c", "d"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(s)
		}()
	}

	wg.Wait()
}
