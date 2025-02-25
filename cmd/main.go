package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello, World!")
	}()
	wg.Wait()
}
