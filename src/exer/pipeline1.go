package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
	/*
	//Counter
	go func() {
		for x := 0; x < 100 ;x++ {
			naturals <- x
		}
		close(naturals)
	}()

	//Squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	//Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}*/
}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for {
		x, ok := <-in
		if !ok {
			break
		}
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}