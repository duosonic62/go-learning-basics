package main

import (
	"fmt"
	"time"
)

func main() {
	// makeChanel()
	// sendChanel()
	// sendGoroutineChanel()
	selectChanel()
}

func makeChanel() {
	ch := make(chan int, 10)
	ch <- 5
	fmt.Println(ch)

	i := <-ch
	fmt.Println(ch)
	fmt.Println(i)
}

func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func sendChanel() {
	ch := make(chan int)
	go receiver(ch)

	i := 0
	for i < 10 {
		ch <- i
		i++
	}
}

func receive(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + " is close.")
}

func sendGoroutineChanel() {
	ch := make(chan int, 20)
	go receive("1st goroutine", ch)
	go receive("2nd goroutine", ch)
	go receive("3rd goroutine", ch)

	i := 0
	for i < 100 {
		ch <- i
		i++
	}
	close(ch)

	time.Sleep(3 * time.Second)
}

func selectChanel() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for {
			i := <-ch1
			ch2 <- (i * 2)
		}
	}()

	go func() {
		for {
			i := <-ch2
			ch3 <- (i - 1)
		}
	}()	

	n := 1
	LOOP:
	for {
		select {
		case ch1 <- n:
			n++
		case i := <-ch2:
			fmt.Println("receive", i)
		case i := <-ch3:
			fmt.Println("receive", i)
		default:
			if n > 100 {
				break LOOP
			}
		}
	}
}