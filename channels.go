package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)
	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 11; i++ {
			c <- i
		}
		close(c)
		q <- 0
	}()
	return c
}

func receive(ch, ch2 <-chan int) {
	for {
		select {
		case v := <-ch:
			fmt.Println("-->", v)
		case <-ch2:
			return
		}
	}
}
