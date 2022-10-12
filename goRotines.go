package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(5)
	go func() {
		fmt.Println("Routine Wait", 1)
		wg.Done()
	}()
	go func() {
		fmt.Println("Routine Wait", 2)
		wg.Done()
	}()
	go func() {
		fmt.Println("Routine Wait", 3)
		wg.Done()
	}()
	go func() {
		fmt.Println("Routine Wait", 4)
		wg.Done()
	}()
	go func() {
		fmt.Println("Routine Wait", 5)
		wg.Done()
	}()
	wg.Wait()
	routine(5)
	wg.Wait()

}

func routine(x int) {
	wg.Add(x)
	for i := 0; i < x; i++ {
		go func(j int) {
			fmt.Println("Routine", j)
			wg.Done()
		}(i)
	}
}
