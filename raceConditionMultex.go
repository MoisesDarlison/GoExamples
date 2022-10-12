package main

import (
	"fmt"
	"runtime"
	"sync"
)

var count int
var mu sync.Mutex
var wg sync.WaitGroup

const quantityGoroutines = 6000

func main() {

	createGORoutines(quantityGoroutines)
	wg.Wait()
	fmt.Println("Total de goroutines:\t", quantityGoroutines, "\nTotal do contador:\t", count)

}

func createGORoutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			mu.Lock()
			v := count
			runtime.Gosched()
			v++
			count = v
			mu.Unlock()
			wg.Done()
		}()
	}
}
